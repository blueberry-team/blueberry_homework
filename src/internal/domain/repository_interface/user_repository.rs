use async_trait::async_trait;
use crate::{
    dto::res::{
        token_user_response::TokenUserResponse,
        user_response::UserResponse,
    },
    internal::domain::entities::user_entity::UserEntity,
};

#[async_trait]
pub trait UserRepository {
    fn new() -> Self where Self: Sized;
    async fn find_by_email(&self, email: String) -> Result<bool, String>;
    async fn find_by_id(&self, id: String) -> Result<bool, String>;
    async fn sign_up(&self, user: UserEntity) -> Result<(), String>;
    async fn log_in(&self, email: String, password: Vec<u8>) -> Result<TokenUserResponse, String>;
    async fn get_user(&self, id: String) -> Result<UserResponse, String>;
}
