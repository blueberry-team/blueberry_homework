use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct CompanyEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub id: Uuid,
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub user_id: Uuid,
    pub company_name: String,
    pub company_address: String,
    pub total_staff: i16,
}

// uuid -> string
fn serialize_uuid<S>(uuid: &Uuid, serializer: S) -> Result<S::Ok, S::Error>
where
    S: serde::Serializer,
{
    serializer.serialize_str(&uuid.to_string())
}

fn deserialize_uuid<'de, D>(deserializer: D) -> Result<Uuid, D::Error>
where
    D: serde::Deserializer<'de>,
{
    let s = String::deserialize(deserializer)?;
    Uuid::parse_str(&s).map_err(serde::de::Error::custom)
}


impl CompanyEntity {
    pub fn new(id: Uuid, user_id: Uuid, company_name: String, company_address: String, total_staff: i16) -> Self {
        Self { id, user_id, company_name, company_address, total_staff }
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct ChangeCompanyEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub user_id: Uuid,
    pub company_name: String,
    pub company_address: String,
    pub total_staff: i16,
}

impl ChangeCompanyEntity {
    pub fn new(user_id: Uuid, company_name: String, company_address: String, total_staff: i16) -> Self {
        Self { user_id, company_name, company_address, total_staff }
    }
}
