use async_trait::async_trait;

use crate::internal::domain::entities::company_entity::CompanyEntity;

#[async_trait]
pub trait CompanyRepository {
    fn new() -> Self where Self: Sized;
    async fn create_company(&self, company: CompanyEntity) -> Result<CompanyEntity, String>;
    async fn get_companies(&self) -> Vec<CompanyEntity>;
}
