use async_trait::async_trait;

use crate::internal::domain::entities::company_entity::CompanyEntity;

#[async_trait]
pub trait CompanyRepository {
    fn new() -> Self where Self: Sized;
    async fn has_company(&self, name: String) -> Result<bool, String>;
    async fn create_company(&self, company: CompanyEntity) -> Result<CompanyEntity, String>;
    async fn get_companies(&self) -> Result<Vec<CompanyEntity>, String>;
}
