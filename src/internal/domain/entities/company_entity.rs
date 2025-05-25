use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::internal::domain::entities::user_entity::{serialize_uuid, deserialize_uuid};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct CompanyEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub id: Uuid,
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub user_id: Uuid,
    pub company_name: String,
    pub company_address: String,
    pub total_staff: u16,
}



impl CompanyEntity {
    pub fn new(id: Uuid, user_id: Uuid, company_name: String, company_address: String, total_staff: u16) -> Self {
        Self { id, user_id, company_name, company_address, total_staff }
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct ChangeCompanyEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub user_id: Uuid,
    pub company_name: String,
    pub company_address: String,
    pub total_staff: u16,
}

impl ChangeCompanyEntity {
    pub fn new(user_id: Uuid, company_name: String, company_address: String, total_staff: u16) -> Self {
        Self { user_id, company_name, company_address, total_staff }
    }
}
