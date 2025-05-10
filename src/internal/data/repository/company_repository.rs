use async_trait::async_trait;
use chrono::Utc;
use futures::TryStreamExt;
use scylla::client::session::Session;
use uuid::Uuid;

use crate::internal::domain::entities::company_entity::CompanyEntity;
use crate::internal::domain::repository_interface::company_repository::CompanyRepository;
use std::sync::{Arc, Mutex};

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
    fn new() -> Self {
        panic!("Use ScyllaCompanyImpl::new(session) instead")
    }

    async fn has_company(&self, name: String) -> Result<bool, String> {
        let query = "SELECT id FROM company WHERE company_name = ?";

        let mut rows = self.session
            .query_iter(query, (name.clone(),))
            .await
            .map_err(|e| format!("Error querying company: {}", e))?
            .rows_stream::<(Uuid,)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        // check result
        let result = rows.try_next().await.map_err(|e| format!("Error getting row: {}", e))?;
        Ok(result.is_some())
    }

    async fn create_company(&self, company: CompanyEntity) -> Result<CompanyEntity, String> {
        // check user already have company
        let check_result = self.has_company(company.name.clone()).await?;
        if check_result {
            return Err(format!("User already have company"));
        }

        let query = "INSERT INTO company (id, name, company_name, create_at, update_at) VALUES (?, ?, ?, ?, ?)";

        let created_timestamp = Utc::now().timestamp();

        let company_data = company.clone();

        self.session
            .query_iter(query, (company_data.id, company_data.name, company_data.company_name, created_timestamp, created_timestamp))
            .await
            .map_err(|e| format!("Error creating company: {}", e))?;

        Ok(company)
    }

    async fn get_companies(&self) -> Result<Vec<CompanyEntity>, String> {
        let query = "SELECT id, name, company_name, create_at, update_at FROM company";

        let rows_stream = self.session
            .query_iter(query, ())
            .await
            .map_err(|e| format!("Error getting row: {}", e))?
            .rows_stream::<(Uuid, String, String, i64, i64)>()
            .map_err(|e| format!("Error getting row: {}", e))?;

        let mut companies = Vec::new();

        let collected_rows: Vec<_> = rows_stream
            .try_collect()
            .await
            .map_err(|e| format!("Error collecting companies: {}", e))?;

        for (id, name, company_name, create_at, update_at) in collected_rows {
            companies.push(CompanyEntity {
                id,
                name,
                company_name,
                create_at,
                update_at,
            });
        }
        Ok(companies)
    }
}
