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
use crate::dto::req::user_req::{ChangeUserReq, GetUserReq, LogInReq, SignUpReq};
use crate::internal::domain::repository_interface::user_repository::UserRepository;
use crate::internal::domain::usecases::user_usecases::change_user_usecase::ChangeUserUsecase;
use crate::internal::domain::usecases::user_usecases::get_user_usecase::GetUserUsecase;
use crate::internal::domain::usecases::user_usecases::sign_up_usecase::SignUpUsecase;
use crate::dto::res::basic_response::BasicResponse;
use crate::internal::domain::usecases::user_usecases::log_in_usecase::LogInUsecase;

pub struct UserHandler;

impl UserHandler {

    pub async fn sign_up_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_req): Json<SignUpReq>,
    ) -> Response<Body> {
        // 유효성 검사 오류 메시지를 동적으로 구성
        if let Err(errors) = user_req.validate() {
            let error_messages = errors.field_errors().iter()
                .flat_map(|(field, errors)| {
                    errors.iter().map(move |err| {
                        format!("{}: {}", field, err.message.clone().unwrap_or_default())
                    })
                })
                .collect::<Vec<String>>()
                .join(", ");

            let response = BasicResponse::bad_request("validation_error".to_string(), error_messages);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = SignUpUsecase::new(repo);

        match usecase.sign_up_usecase(user_req).await {
            Ok(_) => {
                let response = BasicResponse::created("Success".to_string());
                (StatusCode::CREATED, Json(response)).into_response()
            },
            Err(error) => {
                // 에러 유형에 따라 다른 상태 코드 반환
                let status_code = if error.contains("already exists") {
                    StatusCode::CONFLICT
                } else if error.contains("not found") {
                    StatusCode::NOT_FOUND
                } else {
                    StatusCode::BAD_REQUEST
                };

                let response = BasicResponse::bad_request("error".to_string(), error);
                (status_code, Json(response)).into_response()
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

    pub async fn get_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_req): Json<GetUserReq>,
    ) -> Response<Body> {
        // validation for get_user_dto
        if let Err(errors) = user_req.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "name must be 1 and 50 characters".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = GetUserUsecase::new(repo);
        match usecase.get_user_usecase(user_req.id).await {
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

    pub async fn change_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_req): Json<ChangeUserReq>,
    ) -> Response<Body> {
        // validation for change_user_dto
        if let Err(errors) = user_req.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "name must be 1 and 50 characters".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = ChangeUserUsecase::new(repo);
        match usecase.change_user_usecase(user_req).await {
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

