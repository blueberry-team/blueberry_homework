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
        req::user_req::SignUpReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{repository_interface::user_repository::UserRepository, usecases::user_usecases::sign_up_usecase::SignUpUsecase},
};

pub struct SignUpHandler;

impl SignUpHandler {

    pub async fn sign_up_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(user_req): Json<SignUpReq>,
    ) -> Response<Body> {
        // validation error message is dynamic
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
                // each error type has different status code
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
}
