use serde::{Deserialize, Serialize};
use uuid::Uuid;

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct CompanyEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub id: Uuid,
    pub name: String,
    pub company_name: String,
    pub create_at: i64,
    pub update_at: i64,
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
    pub fn new(id: Uuid, name: String, company_name: String, create_at: i64, update_at: i64) -> Self {
        Self { id, name, company_name, create_at, update_at }
    }
}
