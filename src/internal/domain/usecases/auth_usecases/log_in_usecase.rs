use std::sync::Arc;

use crate::{dto::req::user_req::LogInReq, internal::domain::{repository_interface::user_repository::UserRepository, service::password_hash::hash_password}};

pub struct LogInUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl LogInUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { user_repo }
    }

    pub async fn log_in_usecase(&self, log_in_req: LogInReq) -> Result<(), String> {
        let check_user_result = self.user_repo.find_by_email(log_in_req.email.clone()).await?;
        if !check_user_result {
            return Err(format!("User not found"));
        }

        let hashed_password = hash_password(&log_in_req.password).unwrap();

        let user = self.user_repo.log_in(log_in_req.email.clone(), hashed_password).await?;

        // next time, will create token and return it
        println!("user: {:?}", user);

        Ok(())
    }
}
