use std::sync::Arc;
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

    pub async fn create_name(&self, name: String) -> UserEntity {
        self.user_repo.create_name(name).await
    }

    // 여길 만약에 [{name: "NAME"}] 라면 Vec<UserEntity> 로 적용해야하고
    // 현재 구조에선 ["NAME"]로 하는 부분이기때문에 Vec<String> 으로 적용해야함
    pub async fn get_names(&self) -> Vec<UserEntity> {
        self.user_repo.get_names().await
    }

    pub async fn delete_name(&self, index: u32) -> Result<(), String> {
        self.user_repo.delete_name(index).await
    }
}
