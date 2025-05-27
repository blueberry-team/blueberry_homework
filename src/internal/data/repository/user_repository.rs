use std::sync::Arc;
use async_trait::async_trait;
use chrono::{DateTime, Utc};
use scylla::client::session::Session;
use uuid::Uuid;
use crate::dto::res::token::token_user_response::TokenUserResponse;
use crate::dto::res::user_response::UserResponse;
use crate::internal::domain::entities::user_entity::{ChangeUserEntity, UserEntity};
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

    async fn find_by_id(&self, id: Uuid) -> Result<bool, String> {
        let query = "SELECT id FROM user WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (id,))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        Ok(result.is_some())
    }

    async fn find_by_email(&self, email: String) -> Result<bool, String> {
        let query = "SELECT id FROM user WHERE email = ?";

        let mut rows = self.session
            .query_iter(query, (email.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        Ok(result.is_some())
    }

    async fn check_user_role(&self, id: Uuid) -> Result<String, String> {
        let query = "SELECT role FROM user WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (id,))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(String,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((role,)) = result {
            Ok(role)
        } else {
            Err(format!("User not found"))
        }
    }

    async fn sign_up(&self, user: UserEntity) -> Result<(), String> {
        let check_user_result = self.find_by_email(user.email.clone()).await?;
        if check_user_result {
            return Err(format!("User already exists"));
        }

        let create_time = Utc::now().timestamp();

        let query = "INSERT INTO user (id, email, name, password, salt, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)";

        let values = (
            user.id,
            user.email,
            user.name,
            user.password,
            user.salt,
            user.role,
            create_time,
            create_time,
        );

        self.session
            .query_iter(query, values)
            .await
            .map_err(|e| format!("Error signing up user: {}", e))?;

        Ok(())
    }

    async fn get_hashed_password_with_salt(&self, email: String) -> Result<(Vec<u8>, String), String> {
        let query = "SELECT password, salt FROM user WHERE email = ?";

        let mut rows = self.session
            .query_iter(query, (email.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Vec<u8>, String)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((password, salt)) = result {
            Ok((password, salt))
        } else {
            Err(format!("User not found"))
        }
    }

    async fn log_in(&self, email: String, password: Vec<u8>) -> Result<TokenUserResponse, String> {
        let query = "SELECT id, email, name, password FROM user WHERE email = ?";

        let mut rows = self.session
            .query_iter(query, (email.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid, String, String, Vec<u8>,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((id, email, name, stored_password)) = result {
            if stored_password == password {
                Ok(
                    TokenUserResponse {
                    id,
                    email,
                    name,
                }
            )
            } else {
                Err(format!("Invalid password"))
            }
        } else {
            Err(format!("User not found"))
        }
    }

    async fn get_user(&self, id: Uuid) -> Result<UserResponse, String> {
        let query = "SELECT id,email, name, role, created_at, updated_at FROM user WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (id,))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid, String, String, String, i64, i64)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((id, email, name, role, created_at, updated_at)) = result {
            let dt_created_at: DateTime<Utc> = DateTime::from_timestamp(created_at, 0).unwrap();
            let dt_updated_at: DateTime<Utc> = DateTime::from_timestamp(updated_at, 0).unwrap();

            Ok(UserResponse { id, email, name, role, created_at: dt_created_at, updated_at: dt_updated_at })
        } else {
            Err(format!("User not found"))
        }
    }

    async fn change_user(&self, user: ChangeUserEntity) -> Result<(), String> {
        let updated_at = Utc::now().timestamp();

        let query = "UPDATE user SET name = ?, role = ?, updated_at = ? WHERE id = ?";

        self.session
            .query_iter(query, (user.name, user.role, updated_at, user.id))
            .await
            .map_err(|e| format!("Error changing user: {}", e))?;

        Ok(())
    }
}
