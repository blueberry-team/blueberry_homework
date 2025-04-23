use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct UserEntity {
    pub id: String,
    pub name: String,
    pub create_at: DateTime<Utc>,
    pub update_at: DateTime<Utc>,
}

impl UserEntity {
    pub fn new(id: String, name: String, create_at: DateTime<Utc>, update_at: DateTime<Utc>) -> Self {
        Self { id, name, create_at, update_at }
    }
}
