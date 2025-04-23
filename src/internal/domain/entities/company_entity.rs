use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct CompanyEntity {
    pub id: String,
    pub name: String,
    pub company_name: String,
    pub create_at: DateTime<Utc>,
    pub update_at: DateTime<Utc>,
}

impl CompanyEntity {
    pub fn new(id: String, name: String, company_name: String, create_at: DateTime<Utc>, update_at: DateTime<Utc>) -> Self {
        Self { id, name, company_name, create_at, update_at }
    }
}
