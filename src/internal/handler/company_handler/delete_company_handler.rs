use std::sync::Arc;

use axum::{
    body::Body,
    http::{Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};

use crate::{
    dto::res::{basic_response::BasicResponse, token::token_claim::TokenClaim},
    internal::domain::{
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        },
        usecases::company_usecases::delete_company_usecase::DeleteCompanyUsecase,
        }
};

pub struct DeleteCompanyHandler;

impl DeleteCompanyHandler {

    pub async fn delete_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(token_data): Extension<TokenClaim>,
    ) -> Response<Body> {

        let usecase = DeleteCompanyUsecase::new(user_repo, company_repo);
        match usecase.delete_company_usecase(token_data.sub).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
                (StatusCode::OK, Json(response)).into_response()
            }
            Err(error) => {
                let response = BasicResponse::bad_request("error".to_string(), error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }

}
