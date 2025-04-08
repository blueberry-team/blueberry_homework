use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct UserEntity {
    pub name: String,
}

impl UserEntity {
    pub fn new(name: String) -> Self {
        Self { name }
    }
}
