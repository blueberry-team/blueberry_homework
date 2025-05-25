use serde::{Deserialize, Serialize};
use uuid::Uuid;

use crate::dto::req::user_req::UserRole;

#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct UserEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub id: Uuid,
    pub email: String,
    pub name: String,
    #[serde(skip_serializing, skip_deserializing)]
    pub password: Vec<u8>,
    #[serde(skip_serializing, skip_deserializing)]
    pub salt: String,
    pub role: String,
}

// uuid -> string
pub fn serialize_uuid<S>(uuid: &Uuid, serializer: S) -> Result<S::Ok, S::Error>
where
    S: serde::Serializer,
{
    serializer.serialize_str(&uuid.to_string())
}

pub fn deserialize_uuid<'de, D>(deserializer: D) -> Result<Uuid, D::Error>
where
    D: serde::Deserializer<'de>,
{
    let s = String::deserialize(deserializer)?;
    Uuid::parse_str(&s).map_err(serde::de::Error::custom)
}

impl UserEntity {
    pub fn new(id: Uuid, email: String, name: String, password: Vec<u8>, salt: String, role: String) -> Self {
        let role_str = UserRole::from_string(&role).unwrap();
        Self { id, email, name, password, salt, role: role_str.to_string() }
    }
}


#[derive(Debug, Clone, Deserialize, Serialize)]
pub struct ChangeUserEntity {
    #[serde(serialize_with = "serialize_uuid", deserialize_with = "deserialize_uuid")]
    pub id: Uuid,
    pub name: String,
    pub role: String,
}

impl ChangeUserEntity {
    pub fn new(id: Uuid, name: String, role: String) -> Self {
        Self { id, name, role }
    }
}
