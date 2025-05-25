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

    pub async fn change_company_usecase(&self, change_company_req: ChangeCompanyReq, user_id: String) -> Result<(), String> {
        let parsed_user_id =
            Uuid::parse_str(&user_id)
                .map_err(|e| format!("Invalid user id: {}", e))?;

        let check_user = self.user_repo.find_by_id(parsed_user_id).await?;
        if !check_user {
            return Err(format!("User not found"));
        }

        let check_company = self.company_repo.check_company_with_user_id(parsed_user_id).await?;
        if !check_company {
            return Err(format!("Company not found"));
        }

        let company_id = self.company_repo.get_company_with_user_id(parsed_user_id).await?;

        let change_company_entity = ChangeCompanyEntity::new(parsed_user_id, change_company_req.company_name, change_company_req.company_address, change_company_req.total_staff as u16);

        self.company_repo.change_company(change_company_entity, company_id).await?;

        Ok(())
    }
}
