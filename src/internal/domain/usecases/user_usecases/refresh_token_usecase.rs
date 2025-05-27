use std::sync::Arc;

use uuid::Uuid;

use crate::{dto::res::token::token_user_response::TokenUserResponse, internal::domain::{repository_interface::user_repository::UserRepository, utils::jwt::generate_token::generate_token}};

pub struct RefreshTokenUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
    jwt_secret_key: String,
}

impl RefreshTokenUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>, jwt_secret_key: String) -> Self {
        Self { user_repo, jwt_secret_key }
    }

    pub async fn refresh_token_usecase(&self, user_id: String) -> Result<String, String> {
        let parsed_user_id = Uuid::parse_str(&user_id)
            .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user_result = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user_result {
            return Err(format!("User not found"));
        }

        let user = self.user_repo.get_user(parsed_user_id).await?;

        let token_user_response = TokenUserResponse {
            id: user.id,
            email: user.email,
            name: user.name,
        };

        let token = generate_token(self.jwt_secret_key.clone(), token_user_response).await?;

        Ok(token)
    }
}
