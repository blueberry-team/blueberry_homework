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
        usecases::company_usecases::get_company_usecase::GetCompanyUsecase,
    }, utils::jwt::verify_token::verify_token},
};

pub struct GetCompanyHandler;

impl GetCompanyHandler {

    pub async fn get_company_handler(
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

        let usecase = GetCompanyUsecase::new(company_repo, user_repo);
        let company = usecase.get_company_usecase(token_data.sub).await;
        match company {
            Ok(company) => {
                let response = serde_json::json!({
                    "message": "Success",
                    "status_code": StatusCode::OK.as_u16(),
                    "data": company,
                });
                (StatusCode::OK, Json(response)).into_response()
            }
            Err(error) => {
                let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
                println!("{}", error);
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        }
    }

}
