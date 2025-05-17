use std::sync::Arc;
use chrono::Utc;
use uuid::Uuid;

use crate::dto::req::user_req::{ChangeNameReq, UserReq};
use crate::internal::domain::entities::user_entity::UserEntity;
use crate::internal::domain::repository_interface::user_repository::UserRepository;

#[derive(Clone)]
pub struct UserUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl UserUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn create_name_usecase(&self, user_req: UserReq) -> Result<UserEntity, String> {
        let id = Uuid::new_v4();
        let time = Utc::now().timestamp();
        let create_at = time;
        let update_at = time;

        let user = UserEntity::new(id, user_req.name, create_at, update_at);
        self.user_repo.create_name(user).await
    }

    pub async fn change_name_usecase(&self, change_name_req: ChangeNameReq) -> Result<(), String> {
        let formatted_user_id = Uuid::parse_str(&change_name_req.user_id).map_err(|e| format!("Error parsing uuid: {}", e))?;
        self.user_repo.change_name(formatted_user_id, change_name_req.name).await
    }

    // 여길 만약에 [{name: "NAME"}] 라면 Vec<UserEntity> 로 적용해야하고
    // 현재 구조에선 ["NAME"]로 하는 부분이기때문에 Vec<String> 으로 적용해야함
    pub async fn get_names_usecase(&self) -> Result<Vec<UserEntity>, String> {
        self.user_repo.get_all_user_names().await
    }

    pub async fn delete_name_usecase(&self, name: String) -> Result<(), String> {
        self.user_repo.delete_name(name).await
    }
}
