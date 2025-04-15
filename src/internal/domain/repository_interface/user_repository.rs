use async_trait::async_trait;
use crate::internal::domain::entities::user_entity::UserEntity;

#[async_trait]
pub trait UserRepository {
    fn new() -> Self where Self: Sized;
    async fn create_name(&self, user: UserEntity) -> UserEntity;
    async fn get_names(&self) -> Vec<UserEntity>;
    async fn delete_index(&self, index: u32) -> Result<(), String>;
    async fn delete_name(&self, name: String) -> Result<(), String>;
}
