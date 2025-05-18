use std::sync::Arc;

use axum::{
    body::Body,
    http::{Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};
use validator::Validate;

use crate::{
    dto::{
        req::user_req::GetUserReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{
        repository_interface::user_repository::UserRepository,
        usecases::user_usecases::get_user_usecase::GetUserUsecase,
    },
};

pub struct GetUserHandler;

impl GetUserHandler {

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

}
