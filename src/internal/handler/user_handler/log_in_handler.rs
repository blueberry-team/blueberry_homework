use std::sync::Arc;

use axum::{
    extract::State, http::StatusCode, response::IntoResponse, Extension, Json
};
use validator::Validate;

use crate::{
    dto::{
        req::user_req::LogInReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{repository_interface::user_repository::UserRepository, usecases::user_usecases::log_in_usecase::LogInUsecase},
};

pub struct LogInHandler;

impl LogInHandler {

    pub async fn log_in_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        State(jwt_secret_key): State<String>,
        Json(log_in_req): Json<LogInReq>,
    ) -> impl IntoResponse {
        // validation for change_name_dto
        if let Err(errors) = log_in_req.validate() {
            let response = BasicResponse::bad_request("error".to_string(), "name must be 1 and 50 characters".to_string());
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = LogInUsecase::new(repo, jwt_secret_key);

        match usecase.log_in_usecase(log_in_req).await {
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
