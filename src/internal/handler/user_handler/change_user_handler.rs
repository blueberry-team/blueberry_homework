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
        req::user_req::ChangeUserReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{
        repository_interface::user_repository::UserRepository,
        usecases::user_usecases::change_user_usecase::ChangeUserUsecase,
    },
};

pub struct ChangeUserHandler;

impl ChangeUserHandler {

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
