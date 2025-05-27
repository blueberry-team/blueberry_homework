use chrono::Utc;
use jsonwebtoken::{encode, Algorithm, EncodingKey, Header};

use crate::dto::res::token::{token_claim::TokenClaim, token_user_response::TokenUserResponse};

pub async fn generate_token(jwt_secret_key: String, user: TokenUserResponse) -> Result<String, String> {
    // token algorithem
    let chosen_alogorithm = Algorithm::HS512;

    // encoding key from config
    let encoding_key = EncodingKey::from_secret(jwt_secret_key.as_bytes());

    // create timestamp
    let issued_at = Utc::now();

    // create expires at timestamp
    let expires_at = issued_at + chrono::Duration::hours(5);

    // token payload
    let claims = TokenClaim {
        sub: user.id.to_string(),
        email: user.email,
        name: user.name,
        exp: expires_at.timestamp() as usize,
        iat: issued_at.timestamp() as usize,
    };

    let header = Header::new(chosen_alogorithm);

    let token = match encode(&header, &claims, &encoding_key) {
        Ok(token) => token,
        Err(e) => return Err(format!("Failed to encode token: {}", e)),
    };

    Ok(token)
}
