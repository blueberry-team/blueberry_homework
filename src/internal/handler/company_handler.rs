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
    dto::req::company_req::CompanyReq,
    internal::domain::{
        repository_interface::{company_repository::CompanyRepository,
        user_repository::UserRepository},
        usecases::company_usecases::{
            create_company_usecase::CreateCompanyUsecase,
            get_companies_usecase::GetCompaniesUsecase,
        },
    },
    dto::res::basic_response::BasicResponse,
};

pub struct CompanyHandler;

impl CompanyHandler {
    pub async fn create_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(company_req): Json<CompanyReq>,
    ) -> Response<Body> {
        // validation for company_dto
        if let Err(errors) = company_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = CreateCompanyUsecase::new(user_repo, company_repo);

        match usecase.create_company_usecase(company_req).await {
            Ok(_) => {
                let response = BasicResponse::created("Success".to_string());
                (StatusCode::CREATED, Json(response)).into_response()
            }
            Err(error) => {
                let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
                println!("{}", error);
                (StatusCode::BAD_REQUEST, Json(response)).into_response()
            }
        }
    }

    pub async fn get_companies_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
    ) -> Response<Body> {
        let usecase = GetCompaniesUsecase::new(company_repo);
        let companies = usecase.get_companies_usecase().await;

        // if companies is empty return Empty company list please create company
        let (message, status_code) = if companies.is_err() {
            ("Empty company list please create company", StatusCode::OK)
        } else {
            ("Success", StatusCode::OK)
        };

        let response = serde_json::json!({
            "message": message,
            "status_code": status_code.as_u16(),
            "data": companies,
        });

        (StatusCode::OK, Json(response)).into_response()
    }
}
