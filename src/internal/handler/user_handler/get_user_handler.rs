use std::sync::Arc;

use axum::{
    body::Body,
    http::{HeaderMap, Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};

use crate::{
    dto::res::basic_response::BasicResponse,
    internal::{
        domain::{
            repository_interface::user_repository::UserRepository,
            usecases::user_usecases::get_user_usecase::GetUserUsecase,
    },
    utils::jwt::verify_token::verify_token,
    },
};

pub struct GetUserHandler;

impl GetUserHandler {

    pub async fn get_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(jwt_secret_key): Extension<String>,
        headers: HeaderMap,
    ) -> Response<Body> {
        // verify token
        let token_data = match verify_token(jwt_secret_key, &headers).await {
            Ok(token_data) => token_data,
            Err(e) => {
                let response = BasicResponse::bad_request("error".to_string(), e);
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        };

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
