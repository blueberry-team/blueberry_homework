use async_trait::async_trait;
use chrono::{DateTime, Utc};
use crate::internal::domain::entities::user_entity::UserEntity;

#[async_trait]
pub trait UserRepository {
    fn new() -> Self where Self: Sized;
    async fn create_name(&self, user: UserEntity) -> Result<UserEntity, String>;
    async fn change_name(&self, user_id: String, name: String, update_at: DateTime<Utc>) -> Result<(), String>;
    async fn find_by_name(&self, name: String) -> Result<(), String>;
    async fn get_names(&self) -> Vec<UserEntity>;
    async fn delete_index(&self, index: u32) -> Result<(), String>;
    async fn delete_name(&self, name: String) -> Result<(), String>;
}
