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
    #[validate(length(min = 1, max = 50))]
    pub role: String,
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct LogInReq {
    #[validate(email)]
    pub email: String,
    #[validate(length(min = 1, max = 50))]
    pub password: String,
}
