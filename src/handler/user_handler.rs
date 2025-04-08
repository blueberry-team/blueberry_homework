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
use crate::dto::user_dto::UserDto;
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
            let response = BasicResponse::bad_request(format!("Invalid name"));
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
}
