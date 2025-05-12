use std::sync::Arc;

use axum::body::Body;
use axum::http::Response;
use axum::Extension;
use axum::{
    extract::Json,
    http::StatusCode,
    response::IntoResponse,
};
use validator::Validate;
use crate::dto::req::user_req::{SignUpReq, LogInReq};
use crate::internal::domain::repository_interface::user_repository::UserRepository;
use crate::internal::domain::usecases::auth_usecases::sign_up_usecase::SignUpUsecase;
use crate::dto::res::basic_response::BasicResponse;
use crate::internal::domain::usecases::auth_usecases::log_in_usecase::LogInUsecase;

pub struct UserHandler;

impl UserHandler {

    pub async fn sign_up_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_req): Json<SignUpReq>,
    ) -> Response<Body> {
        // validation for user_dto
        if let Err(errors) = user_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = SignUpUsecase::new(repo);

        match usecase.sign_up_usecase(user_req).await {
            Ok(_) => {
                let response = BasicResponse::created("Success".to_string());
                (StatusCode::CREATED, Json(response)).into_response()
            },
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }

    pub async fn log_in_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(log_in_req): Json<LogInReq>,
    ) -> impl IntoResponse {
        // validation for change_name_dto
        if let Err(errors) = log_in_req.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "name must be 1 and 50 characters".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = LogInUsecase::new(repo);
        match usecase.log_in_usecase(log_in_req).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
                (StatusCode::OK, Json(response)).into_response()
            },
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }

}
