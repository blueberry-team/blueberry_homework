use chrono::{DateTime, Utc};
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct UserEntity {
    pub name: String,
    pub create_at: DateTime<Utc>,
}

impl UserEntity {
    pub fn new(name: String, create_at: DateTime<Utc>) -> Self {
        Self { name, create_at }
    }
}
