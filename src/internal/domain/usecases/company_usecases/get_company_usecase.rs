use std::sync::Arc;

use uuid::Uuid;

use crate::{
    dto::res::company_response::CompanyResponse,
    internal::domain::repository_interface::{
        company_repository::CompanyRepository,
        user_repository::UserRepository
    }
};

#[derive(Clone)]
pub struct GetCompanyUsecase {
    company_repo: Arc<dyn CompanyRepository + Send + Sync>,
    user_repo: Arc<dyn UserRepository + Send + Sync>,
}

impl GetCompanyUsecase {
    pub fn new(company_repo: Arc<dyn CompanyRepository + Send + Sync>, user_repo: Arc<dyn UserRepository + Send + Sync>) -> Self {
        Self { company_repo, user_repo }
    }

    pub async fn get_company_usecase(&self, user_id: String) -> Result<CompanyResponse, String> {
        // input id is string type, so need to parse it to uuid type
        let parsed_user_id =
            Uuid::parse_str(&user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        // check user is exist
        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        // check user is boss
        let check_user_role = self.user_repo.check_user_role(parsed_user_id).await?;
        if check_user_role != "boss" {
            return Err(format!("No have permission to get companies"));
        }

        // check company is exist
        let check_company = self.company_repo.check_company_with_user_id(parsed_user_id).await?;
        if !check_company {
            return Err(format!("Company not found"));
        }

        let result = self.company_repo.get_user_company(parsed_user_id).await?;

        Ok(result)
    }
}
