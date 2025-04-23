use async_trait::async_trait;

use crate::internal::domain::entities::company_entity::CompanyEntity;
use crate::internal::domain::repository_interface::company_repository::CompanyRepository;
use std::sync::Mutex;

pub struct CompanyRepositoryImpl {
    companies: std::sync::Mutex<Vec<CompanyEntity>>,
}

impl CompanyRepositoryImpl {
    pub fn new() -> Self {
        Self {
            companies: Mutex::new(Vec::new()),
        }
    }
}

#[async_trait]
impl CompanyRepository for CompanyRepositoryImpl {
    fn new() -> Self {
        CompanyRepositoryImpl::new()
    }

    async fn create_company(&self, company: CompanyEntity) -> Result<CompanyEntity, String> {
        let mut companies = self.companies.lock().unwrap();
        if companies.iter().any(|c| c.name == company.name) {
            return Err(format!("A company with the same value already exists"));
        }
        companies.push(company.clone());
        Ok(company)
    }

    async fn get_companies(&self) -> Vec<CompanyEntity> {
        let companies = self.companies.lock().unwrap();
        companies.clone()
    }
}
