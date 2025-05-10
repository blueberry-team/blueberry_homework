use async_trait::async_trait;
use uuid::Uuid;
use crate::internal::domain::entities::user_entity::UserEntity;

#[async_trait]
pub trait UserRepository {
    fn new() -> Self where Self: Sized;
    async fn create_name(&self, user: UserEntity) -> Result<UserEntity, String>;
    async fn change_name(&self, user_id: Uuid, name: String) -> Result<(), String>;
    async fn find_by_name(&self, name: String) -> Result<bool, String>;
    async fn get_all_user_names(&self) -> Result<Vec<UserEntity>, String>;
    async fn delete_name(&self, name: String) -> Result<(), String>;
}
