use std::sync::Arc;

use chrono::Utc;
use uuid::Uuid;

use crate::{dto::req::user_req::SignUpReq, internal::domain::{entities::UserEntity, repository_interface::user_repository::UserRepository, service::password_hash::hash_password}};

#[derive(Clone)]
pub struct SignUpUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl SignUpUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn sign_up_usecase(&self, sign_up_req: SignUpReq) -> Result<(), String> {
        let id = Uuid::new_v4();
        let created_at = Utc::now().timestamp_millis();
        let hashed_password = hash_password(&sign_up_req.password).unwrap();

        let user = UserEntity::new(
            id,
            sign_up_req.email,
            sign_up_req.name,
            hashed_password,
            sign_up_req.role,
            created_at,
            created_at,
        );

        self.user_repo.sign_up(user).await?;

        Ok(())
    }
}
