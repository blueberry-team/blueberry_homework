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
use crate::dto::user_dto::{DeleteUserDto, UserDto};
use crate::internal::domain::repository_interface::user_repository::UserRepository;
use crate::internal::domain::usecases::user_usecases::UserUsecase;
use crate::res::basic_response::BasicResponse;


pub struct UserHandler;

impl UserHandler {

    pub async fn create_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_dto): Json<UserDto>,
    ) -> Response<Body> {
        // validation for user_dto
        if let Err(errors) = user_dto.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = UserUsecase::new(repo);
        let name = user_dto.name;

        usecase.create_name(name).await;

        let response = BasicResponse::created("Success".to_string());
        (StatusCode::CREATED, Json(response)).into_response()
    }

    pub async fn get_names_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
    ) -> impl IntoResponse {
        let usecase = UserUsecase::new(repo);
        let names = usecase.get_names().await;

        // if names is empty return Empty name list please create name
        let (message, status_code) = if names.is_empty() {
            ("Empty name list please create name", StatusCode::OK)
        } else {
            ("Success", StatusCode::OK)
        };

        let response = serde_json::json!({
            "message": message,
            "status_code": status_code.as_u16(),
            "data": names,
        });

        (StatusCode::OK, Json(response))
    }

    pub async fn delete_name_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        result: Result<Json<DeleteUserDto>, axum::extract::rejection::JsonRejection>,
    ) -> impl IntoResponse {
        // parse json error
        let delete_dto = match result {
            Ok(json) => json.0,
            Err(_) => {
                let response = BasicResponse::bad_request("error".to_string(), "유효한 JSON 형식이 아니거나 인덱스 필드가 누락되었습니다".to_string());
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        };
        
        // validation for delete_dto
        if let Err(errors) = delete_dto.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "인덱스는 0 이상의 값이어야 합니다".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        
        let usecase = UserUsecase::new(repo);
        
        match usecase.delete_name(delete_dto.index).await {
            // if success return ok and message
            Ok(_) => {
                let response = BasicResponse::ok("Successfully deleted name".to_string());
                (StatusCode::OK, Json(response)).into_response()
            },
            // if error return bad_request and error message
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }
}