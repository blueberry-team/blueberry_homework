use std::sync::Arc;

use uuid::Uuid;

use crate::{
    dto::res::user_response::UserResponse,
    internal::domain::repository_interface::user_repository::UserRepository
};

pub struct GetUserUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl GetUserUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn get_user_usecase(&self, user_id: String) -> Result<UserResponse, String> {
        // input id is string type, so need to parse it to uuid type
        let parsed_user_id =
            Uuid::parse_str(&user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        let user = self.user_repo.get_user(parsed_user_id).await?;

        Ok(user)
    }
}
