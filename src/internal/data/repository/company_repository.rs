use async_trait::async_trait;
use chrono::{DateTime, Utc};
use futures::TryStreamExt;
use scylla::client::session::Session;
use uuid::Uuid;

use crate::dto::res::company_response::CompanyResponse;
use crate::internal::domain::entities::company_entity::{ChangeCompanyEntity, CompanyEntity};
use crate::internal::domain::repository_interface::company_repository::CompanyRepository;
use std::sync::Arc;

pub struct CompanyRepositoryImpl {
    session: Arc<Session>,
}

impl CompanyRepositoryImpl {
    pub fn new(session: Arc<Session>) -> Self {
        Self {
            session,
        }
    }
}

#[async_trait]
impl CompanyRepository for CompanyRepositoryImpl {

    async fn check_company_with_user_id(&self, user_id: Uuid) -> Result<bool, String> {
        let query = "SELECT user_id FROM company WHERE user_id = ?";

        let mut rows = self.session
            .query_iter(query, (user_id,))
            .await
            .map_err(|e| format!("Error querying company: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        // check result
        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        Ok(result.is_some())
    }

    async fn get_company_with_user_id(&self, user_id: Uuid) -> Result<Uuid, String> {
        let query = "SELECT id FROM company WHERE user_id = ?";

        let mut rows = self.session
            .query_iter(query, (user_id,))
            .await
            .map_err(|e| format!("Error querying company: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((id,)) = result {
            Ok(id)
        } else {
            Err(format!("Company not found"))
        }
    }

    async fn create_company(&self, company: CompanyEntity) -> Result<(), String> {
        let query = "INSERT INTO company (id, user_id, company_name, company_address, total_staff, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)";

        let created_timestamp = Utc::now().timestamp();

        let values = (
            company.id,
            company.user_id,
            company.company_name,
            company.company_address,
            company.total_staff as i16,
            created_timestamp,
            created_timestamp,
        );

        self.session
            .query_iter(query, values)
            .await
            .map_err(|e| format!("Error creating company: {}", e))?;

        Ok(())
    }

    async fn get_user_company(&self, company_id: Uuid) -> Result<CompanyResponse, String> {
        let query = "SELECT id, user_id, company_name, company_address, total_staff, created_at, updated_at FROM company WHERE id = ?";

        let mut rows = self.session
            .query_iter(query, (company_id,))
            .await
            .map_err(|e| format!("Error querying company: {}", e))?
            .rows_stream::<(Uuid, Uuid, String, String, i16, i64, i64)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        if let Some((id, user_id, company_name, company_address, total_staff, created_at, updated_at)) = result {
            let dt_created_at: DateTime<Utc> = DateTime::from_timestamp(created_at, 0).unwrap();
            let dt_updated_at: DateTime<Utc> = DateTime::from_timestamp(updated_at, 0).unwrap();

            Ok(CompanyResponse { id, user_id, company_name, company_address, total_staff, created_at: dt_created_at, updated_at: dt_updated_at })
        } else {
            Err(format!("Company not found"))
        }
    }

    async fn change_company(&self, company: ChangeCompanyEntity, company_id: Uuid) -> Result<(), String> {
        let query = "UPDATE company SET company_name = ?, company_address = ?, total_staff = ?, updated_at = ? WHERE id = ?";

        let updated_at = Utc::now().timestamp();

        self.session
            .query_iter(query, (company.company_name, company.company_address, company.total_staff as i16, updated_at, company_id))
            .await
            .map_err(|e| format!("Error changing company: {}", e))?;

        Ok(())
    }

    async fn delete_company(&self, user_id: Uuid) -> Result<(), String> {
        let query = "DELETE FROM company WHERE user_id = ?";

        self.session
            .query_iter(query, (user_id,))
            .await
            .map_err(|e| format!("Error deleting company: {}", e))?;

        Ok(())
    }
}
