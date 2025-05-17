use serde::{Deserialize, Serialize};
use validator::Validate;


#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct UserReq {
    #[validate(length(min =1, max = 50))]
    pub name: String,
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct ChangeNameReq {
    #[validate(length(min = 1,))]
    pub user_id: String,
    #[validate(length(min = 1, max = 50))]
    pub name: String,
}

#[derive(Debug, Serialize, Deserialize, Validate)]
pub struct DeleteNameReq {
    #[validate(length(min = 1, max = 50))]
    pub name: String,
}
