use std::sync::Arc;
use async_trait::async_trait;
use chrono::Utc;
use scylla::client::session::Session;
use uuid::Uuid;
use crate::internal::domain::entities::user_entity::UserEntity;
use crate::internal::domain::repository_interface::user_repository::UserRepository;
use futures::TryStreamExt;

// make inmemory list
pub struct UserRepositoryImpl {
    session: Arc<Session>
}

impl UserRepositoryImpl {
    // create user repository
    pub fn new(session: Arc<Session>) -> Self {
        Self {
            session
        }
    }
}

#[async_trait]
impl UserRepository for UserRepositoryImpl {
    fn new() -> Self {
        panic!("Use ScyllaUserImpl::new(session) instead")
    }

    async fn find_by_name(&self, name: String) -> Result<bool, String> {
        let query = "SELECT id FROM user WHERE name = ?";

        let mut rows = self.session
            .query_iter(query, (name.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        // 결과가 있는지 확인
        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        Ok(result.is_some())
    }

    async fn create_name(&self, user: UserEntity) -> Result<UserEntity, String> {
        // check name if user already have same value
        let check_result = self.find_by_name(user.name.clone()).await?;
        if check_result {
            return Err(format!("A name with the same value already exists"));
        }

        let query = "INSERT INTO user (id, name, create_at, update_at) VALUES (?, ?, ?, ?)";

        let user_data = user.clone();

        self.session
            .query_iter(query, (user_data.id, user_data.name, user_data.create_at, user_data.update_at))
            .await
            .map_err(|e| format!("Error creating user: {}", e))?;

        Ok(user)
    }

    async fn change_name(&self, user_id: Uuid, name: String) -> Result<(), String> {
        // check user have in database
        let check_user_query = "SELECT id, name FROM user WHERE id = ?";

        let mut current_user_rows = self.session
            .query_iter(check_user_query, (user_id.clone(),))
            .await
            .map_err(|e| format!("Error checking user: {}", e))?
            .rows_stream::<(Uuid, String)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        if let Some(row_result) = current_user_rows.try_next().await
            .map_err(|e| format!("Error getting row:{}", e))? {
                let (found_id, current_name) = row_result;
                println!("Found user - ID: {}, Name: {}", found_id, current_name);
                if current_name == name {
                    return Err("The name is the same as the current name".into());
                }
            } else {
                return Err(format!("User not found 123: {}", user_id));
            }

        println!("Query result: {:?}", current_user_rows);

        // check name if user already have same value
        let check_user_duplicate_query = "SELECT id FROM user WHERE name = ?";

        let mut conflict_rows_stream = self.session
            .query_iter(check_user_duplicate_query, (name.clone(),))
            .await
            .map_err(|e| format!("Error checking user: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        if conflict_rows_stream.try_next().await
            .map_err(|e| format!("Error getting row: {}", e))?.is_some(){
                return Err(format!("A name with the same value already exists"));
            }

        // update user name
        let updated_timestamp = Utc::now().timestamp();

        let update_query = "UPDATE user SET name = ?, update_at = ? WHERE id = ?";

        self.session
            .query_iter(update_query, (name, updated_timestamp, user_id))
            .await
            .map_err(|e| format!("Error updating user: {}", e))?;

        Ok(())
    }



    async fn get_all_user_names(&self) -> Result<Vec<UserEntity>, String> {
        let query = "SELECT id, name, create_at, update_at FROM user";

        let rows_stream = self.session
            .query_iter(query, ())
            .await
            .map_err(|e| format!("Error getting row: {}", e))?
            .rows_stream::<(Uuid, String, i64, i64)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let mut users = Vec::new();

        // convert stream to vec
        let collected_rows: Vec<_> = rows_stream
            .try_collect()
            .await
            .map_err(|e| format!("Error collecting users: {}", e))?;

        for (id, name, create_at, update_at) in collected_rows {
            users.push(UserEntity {
                id,
                name,
                create_at,
                update_at,
            });
        }

        Ok(users)
    }


    async fn delete_name(&self, name: String) -> Result<(), String> {
        // check user exist
        let check_query = "SELECT id FROM user WHERE name = ?";
        let mut id_stream = self.session
            .query_iter(check_query, (name.clone(),))
            .await
            .map_err(|e| format!("Error checking user: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error creating stream: {}", e))?;

        // find id
        if let Some((user_id,)) = id_stream.try_next().await
            .map_err(|e| format!("Error getting row: {}", e))? {

                // delete user using id
                let delete_query = "DELETE FROM user WHERE id = ?";
                self.session
                    .query_iter(delete_query, (user_id,))
                    .await
                    .map_err(|e| format!("Error deleting user: {}", e))?;

                Ok(())
            } else {
                Err(format!("User not found: {}", name))
            }
    }
}
