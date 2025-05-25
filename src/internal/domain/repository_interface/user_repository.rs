use async_trait::async_trait;
use uuid::Uuid;
use crate::{
    dto::{
        res::{
            token::token_user_response::TokenUserResponse,
            user_response::UserResponse,
        },
    },
    internal::domain::entities::user_entity::{ChangeUserEntity, UserEntity},
};

#[async_trait]
pub trait UserRepository {
    fn new() -> Self where Self: Sized;
    async fn find_by_email(&self, email: String) -> Result<bool, String>;
    async fn find_by_id(&self, id: Uuid) -> Result<bool, String>;
    async fn check_user_role(&self, id: Uuid) -> Result<String, String>;
    async fn sign_up(&self, user: UserEntity) -> Result<(), String>;
    async fn get_hashed_password_with_salt(&self, email: String) -> Result<(Vec<u8>, String), String>;
    async fn log_in(&self, email: String, password: Vec<u8>) -> Result<TokenUserResponse, String>;
    async fn get_user(&self, id: Uuid) -> Result<UserResponse, String>;
    async fn change_user(&self, user: ChangeUserEntity) -> Result<(), String>;
}
