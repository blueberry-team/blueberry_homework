use std::sync::Arc;

use axum::{
    body::Body,
    http::{HeaderMap, Response, StatusCode},
    response::IntoResponse,
    Extension,
    Json,
};

use crate::{
    dto::{
        res::basic_response::BasicResponse,
    },
    internal::{domain::{
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        },
        usecases::company_usecases::delete_company_usecase::DeleteCompanyUsecase,
    }, utils::jwt::verify_token::verify_token},
};

pub struct DeleteCompanyHandler;

impl DeleteCompanyHandler {

    pub async fn delete_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Extension(jwt_secret_key): Extension<String>,
        headers: HeaderMap,
    ) -> Response<Body> {

        // verify token
        let token_data = match verify_token(jwt_secret_key, &headers).await {
            Ok(token_data) => token_data,
            Err(e) => {
                let response = BasicResponse::bad_request("error".to_string(), e);
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        };

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
