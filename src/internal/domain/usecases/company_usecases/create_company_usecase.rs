use std::sync::Arc;
use chrono::Utc;
use uuid::Uuid;

use crate::{dto::req::company_req::CompanyReq, internal::domain::{entities::CompanyEntity, repository_interface::{company_repository::CompanyRepository, user_repository::UserRepository}}};

#[derive(Clone)]
pub struct CreateCompanyUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
    company_repo: Arc<dyn CompanyRepository + Send + Sync>,
}

impl CreateCompanyUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>, company_repo: Arc<dyn CompanyRepository + Send + Sync>) -> Self {
        Self { user_repo, company_repo }
    }

    pub async fn create_company_usecase(&self, company_req: CompanyReq) -> Result<CompanyEntity, String> {
        let time = Utc::now().timestamp();
        let id = Uuid::new_v4();
        let create_time = time;
        let update_time = time;
        let check_user = self.user_repo.find_by_name(company_req.name.clone()).await;
        if check_user.is_err() {
            return Err(check_user.unwrap_err());
        }
        let company = CompanyEntity::new(id, company_req.name, company_req.company_name, create_time, update_time);
        self.company_repo.create_company(company).await
    }
}
