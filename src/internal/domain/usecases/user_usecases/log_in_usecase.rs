use std::sync::Arc;

use crate::{
    dto::{
        req::user_req::LogInReq,
    }, internal::domain::{
        repository_interface::user_repository::UserRepository,
        service::password_hash::verify_password, utils::jwt::generate_token::generate_token
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

        let token = generate_token(self.jwt_secret_key.clone(), user).await?;

        Ok(token)
    }

}
