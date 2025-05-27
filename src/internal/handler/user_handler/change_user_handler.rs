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
        res::{
            basic_response::BasicResponse,
            token::token_claim::TokenClaim
        },
    },
    internal::{
        domain::{
        repository_interface::user_repository::UserRepository,
        usecases::user_usecases::change_user_usecase::ChangeUserUsecase,
        },
    },
};

pub struct ChangeUserHandler;

impl ChangeUserHandler {

    pub async fn change_user_handler(
        Extension(repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(token_data): Extension<TokenClaim>,
        Json(change_user_req): Json<ChangeUserReq>,
    ) -> Response<Body> {
        // validation for change_user_dto
        if let Err(errors) = change_user_req.validate() {
            let error_messages = errors.field_errors().iter()
                .flat_map(|(field, errors)| {
                    errors.iter().map(move |err| {
                        format!("{}: {}", field, err.message.clone().unwrap_or_default())
                    })
                })
                .collect::<Vec<String>>()
                .join(", ");

            let response = BasicResponse::bad_request(format!("error"), error_messages);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }

        let usecase = ChangeUserUsecase::new(repo);

        match usecase.change_user_usecase(token_data.sub, change_user_req).await {
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
