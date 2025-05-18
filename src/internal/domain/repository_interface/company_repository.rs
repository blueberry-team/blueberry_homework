use async_trait::async_trait;
use uuid::Uuid;

use crate::{
    dto::res::company_response::CompanyResponse,
    internal::domain::entities::company_entity::{
        ChangeCompanyEntity,
        CompanyEntity
    }
};

#[async_trait]
pub trait CompanyRepository {
    fn new() -> Self where Self: Sized;
    async fn check_company_with_user_id(&self, user_id: Uuid) -> Result<bool, String>;
    async fn create_company(&self, company: CompanyEntity) -> Result<(), String>;
    async fn get_user_company(&self, user_id: Uuid) -> Result<CompanyResponse, String>;
    async fn change_company(&self, company: ChangeCompanyEntity) -> Result<(), String>;
    async fn delete_company(&self, user_id: Uuid) -> Result<(), String>;
}
