use axum::{
    body::Body, extract::{
        Request, State
    }, http::{HeaderMap, StatusCode}, middleware::Next, response::{IntoResponse, Response}, Json
};
use jsonwebtoken::{decode, Algorithm, DecodingKey, Validation};
use chrono::Utc;
use crate::dto::res::{basic_response::BasicResponse, token::token_claim::TokenClaim};

pub async fn verify_token_middleware(
    State(jwt_secret_key): State<String>,
    mut req: Request,
    next: Next,
) -> Result<Response<Body>, (StatusCode, Response<Body>)> {

    let headers = req.headers();

    // verify token
    let (token_data, should_refresh) = match verify_token(jwt_secret_key, headers).await {
        Ok((token_data, should_refresh)) => (token_data, should_refresh),
        Err(e) => {
            let response = BasicResponse::bad_request("error".to_string(), e);
            return Err((StatusCode::UNAUTHORIZED, Json(response).into_response()));
        }
    };

    // add token data to request
    req.extensions_mut().insert(token_data);
    req.extensions_mut().insert(should_refresh);

    // call next middleware
    let res = next.run(req).await;

    Ok(res)
}

// function to verify token
async fn verify_token(jwt_secret_key: String, headers: &HeaderMap) -> Result<(TokenClaim, bool), String> {

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
    validation.validate_exp = true;

    // decode token
    let token_data = match decode::<TokenClaim>(&get_token_from_header, &decoding_key, &validation) {
        Ok(token_data) => token_data,
        Err(e) => {
            return Err(e.to_string());
        }
    };

    // check token expiration
    let current_time = Utc::now().timestamp() as usize;
    let time_until_expiry = token_data.claims.exp.saturating_sub(current_time);
    let should_refresh_token = time_until_expiry < 3600;


    Ok((token_data.claims, should_refresh_token))
}

// function to get token from header
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
