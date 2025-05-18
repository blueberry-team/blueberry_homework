use std::sync::Arc;

use uuid::Uuid;

use crate::{
    dto::req::company_req::ChangeCompanyReq,
    internal::domain::{
        entities::company_entity::ChangeCompanyEntity,
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        }
    }
};

#[derive(Clone)]
pub struct ChangeCompanyUsecase {
    user_repo: Arc<dyn UserRepository + Send + Sync>,
    company_repo: Arc<dyn CompanyRepository + Send + Sync>,
}

impl ChangeCompanyUsecase {
    pub fn new(user_repo: Arc<dyn UserRepository + Send + Sync>, company_repo: Arc<dyn CompanyRepository + Send + Sync>) -> Self {
        Self { user_repo, company_repo }
    }

    pub async fn change_company_usecase(&self, change_company_req: ChangeCompanyReq) -> Result<(), String> {
        let parsed_user_id =
            Uuid::parse_str(&change_company_req.user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        let check_company = self.company_repo.check_company_with_user_id(parsed_user_id).await?;
        if !check_company {
            return Err(format!("Company not found"));
        }

        let change_company_entity = ChangeCompanyEntity::new(parsed_user_id, change_company_req.company_name, change_company_req.company_address, change_company_req.company_phone, change_company_req.total_staff);

        self.company_repo.change_company(change_company_entity).await?;

        Ok(())
    }
}
