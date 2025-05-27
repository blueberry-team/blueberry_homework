use std::sync::Arc;

use axum::{
    body::Body, extract::State, http::{Response, StatusCode}, response::IntoResponse, Extension, Json
};

use crate::{
    dto::res::{basic_response::BasicResponse, token::token_claim::TokenClaim},
    internal::{
        domain::{
            repository_interface::user_repository::UserRepository,
            usecases::user_usecases::refresh_token_usecase::RefreshTokenUsecase,
        },
    },
};

pub struct RefreshTokenHandler;

impl RefreshTokenHandler {
    pub async fn refresh_token_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(token_data): Extension<TokenClaim>,
        State(jwt_secret_key): State<String>,
    ) -> Response<Body> {
        let usecase = RefreshTokenUsecase::new(repo, jwt_secret_key);

        match usecase.refresh_token_usecase(token_data.sub).await {
            Ok(token) => {
                let response = serde_json::json!({
                    "message": "Success",
                    "status_code": StatusCode::OK.as_u16(),
                    "data": token,
                });
                (StatusCode::OK, Json(response)).into_response()
            },
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }
}
