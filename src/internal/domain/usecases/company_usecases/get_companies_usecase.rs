use std::sync::Arc;

use crate::internal::domain::{
    entities::CompanyEntity,
    repository_interface::company_repository::CompanyRepository,
};

#[derive(Clone)]
pub struct GetCompaniesUsecase {
    company_repo: Arc<dyn CompanyRepository + Send + Sync>,
}

impl GetCompaniesUsecase {
    pub fn new(company_repo: Arc<dyn CompanyRepository + Send + Sync>) -> Self {
        Self { company_repo }
    }

    pub async fn get_companies_usecase(&self) -> Vec<CompanyEntity> {
        self.company_repo.get_companies().await
    }
}
