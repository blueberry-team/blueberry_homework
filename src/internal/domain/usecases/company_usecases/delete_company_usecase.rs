use std::sync::Arc;

use uuid::Uuid;

use crate::{dto::req::company_req::GetCompanyReq, internal::domain::repository_interface::{
    company_repository::CompanyRepository,
    user_repository::UserRepository
}};

#[derive(Clone)]
pub struct DeleteCompanyUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
    company_repo: Arc<dyn CompanyRepository + Send + Sync>,
}

impl DeleteCompanyUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>, company_repo: Arc<dyn CompanyRepository + Send + Sync>) -> Self {
        Self { user_repo, company_repo }
    }

    pub async fn delete_company_usecase(&self, get_company_req: GetCompanyReq) -> Result<(), String> {
        let parsed_user_id =
            Uuid::parse_str(&get_company_req.user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        let check_company = self.company_repo.check_company_with_user_id(parsed_user_id).await?;
        if !check_company {
            return Err(format!("Company not found"));
        }

        let check_user_role = self.user_repo.check_user_role(parsed_user_id).await?;
        if check_user_role != "boss" {
            return Err(format!("No have permission to delete company"));
        }

        self.company_repo.delete_company(parsed_user_id).await?;

        Ok(())
    }
}

