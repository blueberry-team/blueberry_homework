use std::sync::Arc;

use axum::{
    body::Body,
    http::{HeaderMap, Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};
use validator::Validate;

use crate::{
    dto::{
        req::company_req::ChangeCompanyReq,
        res::basic_response::BasicResponse,
    },
    internal::{domain::{
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        },
        usecases::company_usecases::change_company_usecase::ChangeCompanyUsecase,
    }, utils::jwt::verify_token::verify_token},
};

pub struct ChangeCompanyHandler;

impl ChangeCompanyHandler {

    pub async fn change_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(jwt_secret_key): Extension<String>,
        headers: HeaderMap,
        Json(change_company_req): Json<ChangeCompanyReq>,
    ) -> Response<Body> {
        // validation for change_company_dto
        if let Err(errors) = change_company_req.validate() {
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

        // verify token
        let token_data = match verify_token(jwt_secret_key, &headers).await {
            Ok(token_data) => token_data,
            Err(e) => {
                let response = BasicResponse::bad_request("error".to_string(), e);
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        };

        let usecase = ChangeCompanyUsecase::new(user_repo, company_repo);
        match usecase.change_company_usecase(change_company_req, token_data.sub).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
                (StatusCode::OK, Json(response)).into_response()
            }
            Err(error) => {
                let response = BasicResponse::bad_request(format!("error"), format!("{}", error));
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        }
    }

}
