use std::sync::Arc;

use chrono::Utc;
use jsonwebtoken::{encode, Algorithm, EncodingKey, Header};

use crate::{
    dto::{
        req::user_req::LogInReq,
        res::token::{token_claim::TokenClaim, token_user_response::TokenUserResponse}
    }, internal::domain::{
        repository_interface::user_repository::UserRepository,
        service::password_hash::verify_password
    }
};

pub struct LogInUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
    jwt_secret_key: String,
}

impl LogInUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>, jwt_secret_key: String) -> Self {
        Self { user_repo, jwt_secret_key }
    }

    pub async fn log_in_usecase(&self, log_in_req: LogInReq) -> Result<String, String> {
        let check_user_result = self.user_repo.find_by_email(log_in_req.email.clone()).await?;
        if !check_user_result {
            return Err(format!("User not found"));
        }

        let hashed_password_with_salt = self.user_repo.get_hashed_password_with_salt(log_in_req.email.clone()).await?;

        let verify_password_result = verify_password(&log_in_req.password, &hashed_password_with_salt.0, &hashed_password_with_salt.1);

        if !verify_password_result.unwrap() {
            return Err(format!("Invalid password"));
        }

        let user = self.user_repo.log_in(log_in_req.email.clone(), hashed_password_with_salt.0).await?;

        let token = self.generate_token(user).await?;

        Ok(token)
    }

    async fn generate_token(&self, user: TokenUserResponse) -> Result<String, String> {
        // token algorithem
        let chosen_alogorithm = Algorithm::HS512;

        // encoding key from config
        let encoding_key = EncodingKey::from_secret(self.jwt_secret_key.as_bytes());

        // create timestamp
        let issued_at = Utc::now();

        // create expires at timestamp
        let expires_at = issued_at + chrono::Duration::days(1);

        // token payload
        let claims = TokenClaim {
            sub: user.id.to_string(),
            email: user.email,
            name: user.name,
            exp: expires_at.timestamp() as usize,
            iat: issued_at.timestamp() as usize,
        };

        let header = Header::new(chosen_alogorithm);

        let token = match encode(&header, &claims, &encoding_key) {
            Ok(token) => token,
            Err(e) => return Err(format!("Failed to encode token: {}", e)),
        };

        Ok(token)
    }
}
