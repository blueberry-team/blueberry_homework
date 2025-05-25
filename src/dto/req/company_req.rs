use serde::{Deserialize, Serialize};
use validator::Validate;

#[derive(Debug, Serialize, Deserialize, Clone, Validate)]
pub struct CreateCompanyReq {
    #[validate(length(min = 1, max = 50))]
    pub company_name: String,
    #[validate(length(min = 1, max = 50))]
    pub company_address: String,
    #[validate(length(min = 1, max = 50))]
    pub total_staff: String,
}

#[derive(Debug, Serialize, Deserialize, Clone, Validate)]
pub struct ChangeCompanyReq {
    #[validate(length(min = 1, max = 50))]
    pub company_name: String,
    #[validate(length(min = 1, max = 50))]
    pub company_address: String,
    #[validate(range(min = 1, max = 10000))]
    pub total_staff: i16,
}
