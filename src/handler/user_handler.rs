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
use crate::dto::user_dto::{ChangeNameDto, DeleteNameDto, DeleteNameIndexDto, UserDto};
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

        match usecase.create_name_usecase(user_dto).await {
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

    pub async fn change_name_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(change_name_dto): Json<ChangeNameDto>,
    ) -> impl IntoResponse {
        // validation for change_name_dto
        if let Err(errors) = change_name_dto.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "name must be 1 and 50 characters".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = UserUsecase::new(repo);
        match usecase.change_name_usecase(change_name_dto).await {
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

    pub async fn get_names_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
    ) -> impl IntoResponse {
        let usecase = UserUsecase::new(repo);
        let names = usecase.get_names_usecase().await;

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

    pub async fn delete_index_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        result: Result<Json<DeleteNameIndexDto>, axum::extract::rejection::JsonRejection>,
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

        match usecase.delete_index_usecase(delete_dto.index).await {
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

    pub async fn delete_name_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(delete_name_dto): Json<DeleteNameDto>,
    ) -> impl IntoResponse {

        // validation for delete_name_dto
        if let Err(errors) = delete_name_dto.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "이름은 1자 이상 50자 이하여야 합니다".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = UserUsecase::new(repo);
        match usecase.delete_name_usecase(delete_name_dto.name).await {
            Ok(_) => {
                let response = BasicResponse::ok("Successfully deleted name".to_string());
                (StatusCode::OK, Json(response)).into_response()
            },
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }

}
