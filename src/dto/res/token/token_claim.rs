use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize)]
pub struct TokenClaim {
    pub sub: String,
    pub email: String,
    pub name: String,
    pub exp: usize,
    pub iat: usize,
}
