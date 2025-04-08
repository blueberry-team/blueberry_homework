use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct BasicResponse {
    pub message: String,
}

impl BasicResponse {
    pub fn new(message: String) -> Self {
        Self {
            message,
        }
    }

    pub fn ok(message: String) -> Self {
        Self {
            message,
        }
    }

    pub fn created(message: String) -> Self {
        Self {
            message,
        }
    }

    pub fn bad_request(message: String) -> Self {
        Self {
            message,
        }
    }
}
