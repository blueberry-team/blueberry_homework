use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct BasicResponse {
    pub message: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub error: Option<String>,
}

impl BasicResponse {
    pub fn new(message: String) -> Self {
        Self {
            message,
            error: None,
        }
    }

    pub fn ok(message: String) -> Self {
        Self {
            message,
            error: None,
        }
    }

    pub fn created(message: String) -> Self {
        Self {
            message,
            error: None,
        }
    }

    pub fn bad_request(message: String, error: String) -> Self {
        Self {
            message,
            error: Some(error),
        }
    }
}
