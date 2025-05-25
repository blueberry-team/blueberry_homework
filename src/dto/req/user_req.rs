use serde::{Deserialize, Serialize};
use validator::Validate;

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct SignUpReq {
    #[validate(email)]
    pub email: String,
    #[validate(length(min =1, max = 50))]
    pub name: String,
    #[validate(length(min = 1, max = 50))]
    pub password: String,
    pub role: String,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(rename_all = "lowercase")] // enum 값을 소문자로 직렬화
pub enum UserRole {
    Worker,
    Boss,
}

impl UserRole {
    // validate user role
    pub fn from_string(s: &str) -> Result<Self, String> {
        match s.to_lowercase().as_str() {
            "worker" => Ok(UserRole::Worker),
            "boss" => Ok(UserRole::Boss),
            _ => Err(format!("Invalid user role: {}", s)),
        }
    }

    pub fn to_string(&self) -> String {
        match self {
            UserRole::Worker => "worker".to_string(),
            UserRole::Boss => "boss".to_string(),
        }
    }
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct LogInReq {
    #[validate(email)]
    pub email: String,
    #[validate(length(min = 1, max = 50))]
    pub password: String,
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct ChangeUserReq {
    #[validate(length(min = 1, max = 50))]
    pub name: String,
    #[validate(length(min = 1, max = 50))]
    pub role: String,
}
