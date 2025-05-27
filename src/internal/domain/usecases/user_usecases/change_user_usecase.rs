use std::sync::Arc;

use uuid::Uuid;

use crate::{
    dto::req::user_req::ChangeUserReq,
    internal::domain::{
        entities::user_entity::ChangeUserEntity,
        repository_interface::user_repository::UserRepository
    }
};

pub struct ChangeUserUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl ChangeUserUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn change_user_usecase(&self, user_id: String, user_req: ChangeUserReq) -> Result<(), String> {
        let parsed_user_id =
            Uuid::parse_str(&user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user = self.user_repo.find_by_id(parsed_user_id.clone()).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        let change_user_entity = ChangeUserEntity::new(parsed_user_id, user_req.name, user_req.role);

        self.user_repo.change_user(change_user_entity).await?;

        Ok(())
    }
}
