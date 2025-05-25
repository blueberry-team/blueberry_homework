use axum::http::HeaderMap;
use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};

use crate::dto::res::token::token_claim::TokenClaim;

pub async fn verify_token(jwt_secret_key: String, headers: &HeaderMap) -> Result<TokenClaim, String> {

    // get token from header
    let get_token_from_header = match get_header_token(headers).await {
        Ok(token) => token,
        Err(e) => {
            return Err(e);
        }
    };

    // decode token
    let chosen_algorithm = Algorithm::HS512;

    // get decoding key
    let decoding_key = DecodingKey::from_secret(jwt_secret_key.as_bytes());

    // set validation
    let mut validation = Validation::new(chosen_algorithm);
    validation.validate_exp = false;

    // decode token
    let token_data = match decode::<TokenClaim>(&get_token_from_header, &decoding_key, &validation) {
        Ok(token_data) => token_data,
        Err(e) => {
            return Err(e.to_string());
        }
    };

    Ok(token_data.claims)
}

async fn get_header_token(headers: &HeaderMap) -> Result<String, String> {
    // get token from header
    let token = match headers.get("Authorization") {
        Some(header_value) => {
            let auth_str = header_value.to_str().map_err(|_| "invalid header value")?;
            if auth_str.starts_with("Bearer ") {
                auth_str[7..].to_string()
            } else {
                return Err("invalid header value".to_string());
            }
        }
        None => {
            return Err("no token provided".to_string());
        }
    };

    Ok(token)
}
