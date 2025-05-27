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
        req::company_req::CreateCompanyReq,
        res::{basic_response::BasicResponse, token::token_claim::TokenClaim},
    },
    internal::{
        domain::{
            repository_interface::{
                company_repository::CompanyRepository,
                user_repository::UserRepository
            },
            usecases::company_usecases::create_company_usecase::CreateCompanyUsecase,
        },
    },
};

pub struct CreateCompanyHandler;

impl CreateCompanyHandler {

    pub async fn create_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(token_data): Extension<TokenClaim>,
        Json(company_req): Json<CreateCompanyReq>,
    ) -> Response<Body> {
        // validation for company_dto
        if let Err(errors) = company_req.validate() {
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


        let usecase = CreateCompanyUsecase::new(user_repo, company_repo);

        match usecase.create_company_usecase(company_req, token_data.sub).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
                (StatusCode::OK, Json(response)).into_response()
            }
            Err(error) => {
                // each error type has different status code
                let status_code = if error.contains("already exists") {
                    StatusCode::CONFLICT
                } else if error.contains("not found") {
                    StatusCode::NOT_FOUND
                } else {
                    StatusCode::BAD_REQUEST
                };

                let response = BasicResponse::bad_request(format!("error"), error);
                (status_code, Json(response)).into_response()
            }
        }
    }

}
