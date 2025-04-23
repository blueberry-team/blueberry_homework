use std::sync::Arc;
use chrono::Utc;
use uuid::Uuid;

use crate::dto::user_dto::{ChangeNameDto, UserDto};
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

    pub async fn create_name_usecase(&self, user_dto: UserDto) -> Result<UserEntity, String> {
        let time = Utc::now();
        let id = Uuid::new_v4().to_string();
        let create_time = time;
        let update_time = time;
        let user = UserEntity::new(id, user_dto.name, create_time, update_time);
        self.user_repo.create_name(user).await
    }

    pub async fn change_name_usecase(&self, change_name_dto: ChangeNameDto) -> Result<(), String> {
        let update_time = Utc::now();
        self.user_repo.change_name(change_name_dto.user_id, change_name_dto.name, update_time).await
    }

    // 여길 만약에 [{name: "NAME"}] 라면 Vec<UserEntity> 로 적용해야하고
    // 현재 구조에선 ["NAME"]로 하는 부분이기때문에 Vec<String> 으로 적용해야함
    pub async fn get_names_usecase(&self) -> Vec<UserEntity> {
        self.user_repo.get_names().await
    }

    pub async fn delete_index_usecase(&self, index: u32) -> Result<(), String> {
        self.user_repo.delete_index(index).await
    }

    pub async fn delete_name_usecase(&self, name: String) -> Result<(), String> {
        self.user_repo.delete_name(name).await
    }
}
