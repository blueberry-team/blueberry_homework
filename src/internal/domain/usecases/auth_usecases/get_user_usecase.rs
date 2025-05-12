use std::sync::Arc;

use crate::{dto::res::user_response::UserResponse, internal::domain::repository_interface::user_repository::UserRepository};

pub struct  GetUserUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl GetUserUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn get_user_usecase(&self, user_id: String) -> Result<UserResponse, String> {
        let check_user = self.user_repo.find_by_id(user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        

        Ok(user)
    }
}