use std::sync::Arc;

use axum::{
    body::Body,
    http::{Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};

use crate::{
    dto::res::{basic_response::BasicResponse, token::token_claim::TokenClaim},
    internal::{
        domain::{
            repository_interface::user_repository::UserRepository,
            usecases::user_usecases::get_user_usecase::GetUserUsecase,
        },
    },
};

pub struct GetUserHandler;

impl GetUserHandler {

    pub async fn get_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(token_data): Extension<TokenClaim>,
    ) -> Response<Body> {

        let usecase = GetUserUsecase::new(repo);

        match usecase.get_user_usecase(token_data.sub).await {
            Ok(user) => {
                let response = serde_json::json!({
                    "message": "Success",
                    "status_code": StatusCode::OK.as_u16(),
                    "data": user,
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
