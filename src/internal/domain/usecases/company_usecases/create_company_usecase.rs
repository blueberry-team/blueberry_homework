use std::{sync::Arc, str::FromStr};
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

    pub async fn create_company_usecase(&self, company_req: CompanyReq) -> Result<(), String> {
        let id = Uuid::new_v4();
        let parsed_user_id =
            Uuid::parse_str(&company_req.user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let parsed_total_staff =
            i16::from_str(&company_req.total_staff.to_string())
                .map_err(|e| format!("Invalid total staff: {}", e))?;

        // check user is exist
        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        // check user is boss
        let check_user_role = self.user_repo.check_user_role(parsed_user_id).await?;
        if check_user_role != "boss" {
            return Err(format!("No have permission to create company"));
        }

        // check company is exist
        let check_company = self.company_repo.check_company_with_user_id(parsed_user_id).await?;
        if check_company {
            return Err(format!("Company already exist"));
        }

        let company = CompanyEntity::new(id, parsed_user_id, company_req.company_name, company_req.company_address, parsed_total_staff);
        self.company_repo.create_company(company).await?;

        Ok(())
    }
}
