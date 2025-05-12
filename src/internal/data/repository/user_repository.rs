use std::sync::Arc;
use async_trait::async_trait;
use chrono::{DateTime, Utc};
use scylla::client::session::Session;
use uuid::Uuid;
use crate::dto::res::token_user_response::TokenUserResponse;
use crate::dto::res::user_response::UserResponse;
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

    async fn find_by_id(&self, id: String) -> Result<bool, String> {
        let query = "SELECT id FROM user WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (id.clone(),))
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

    async fn sign_up(&self, user: UserEntity) -> Result<(), String> {
        let check_user_result = self.find_by_email(user.email.clone()).await?;
        if check_user_result {
            return Err(format!("User already exists"));
        }

        let query = "INSERT INTO user (id, email, name, password, role, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)";

        let values = (
            user.id,
            user.email,
            user.name,
            user.password,
            user.role,
            user.created_at,
            user.updated_at,
        );

        self.session
            .query_iter(query, values)
            .await
            .map_err(|e| format!("Error signing up user: {}", e))?;

        Ok(())
    }

    async fn log_in(&self, email: String, password: Vec<u8>) -> Result<TokenUserResponse, String> {
        let query = "SELECT id, email, name, password, role FROM user WHERE email = ?";

        let mut rows = self.session
            .query_iter(query, (email.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(Uuid, String, String, Vec<u8>, String)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((id, email, name, stored_password, role)) = result {
            if stored_password == password {
                Ok(
                    TokenUserResponse {
                    id,
                    email,
                    name,
                    role,
                    created_at: 0,
                    updated_at: 0,
                }
            )
            } else {
                Err(format!("Invalid password"))
            }
        } else {
            Err(format!("User not found"))
        }
    }

    async fn get_user(&self, id: String) -> Result<UserResponse, String> {
        let query = "SELECT email, name, role, created_at, updated_at FROM user WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (id.clone(),))
            .await
            .map_err(|e| format!("Error querying user: {}", e))?
            .rows_stream::<(String, String, String, i64, i64)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((email, name, role, created_at, updated_at)) = result {
            let dt_created_at: DateTime<Utc> = DateTime::from_timestamp(created_at, 0).unwrap();
            let dt_updated_at: DateTime<Utc> = DateTime::from_timestamp(updated_at, 0).unwrap();

            Ok(UserResponse { email, name, role, created_at: dt_created_at, updated_at: dt_updated_at })
        } else {
            Err(format!("User not found"))
        }
    }
}
