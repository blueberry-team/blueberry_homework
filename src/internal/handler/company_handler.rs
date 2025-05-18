use std::sync::Arc;

use axum::{
    body::Body,
    http::{
        Response,
        StatusCode
    },
    response::IntoResponse,
    Extension,
    Json,
};
use validator::Validate;

use crate::{
    dto::{req::company_req::{ChangeCompanyReq, CompanyReq, GetCompanyReq}, res::basic_response::BasicResponse},
    internal::domain::{
        repository_interface::{company_repository::CompanyRepository,
        user_repository::UserRepository},
        usecases::company_usecases::{
            change_company_usecase::ChangeCompanyUsecase, create_company_usecase::CreateCompanyUsecase, delete_company_usecase::DeleteCompanyUsecase, get_company_usecase::GetCompanyUsecase
        },
    },
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

        match usecase.create_company_usecase(company_req).await {
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

    pub async fn get_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(get_company_req): Json<GetCompanyReq>,
    ) -> Response<Body> {
        // validation for get_company_dto
        if let Err(errors) = get_company_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = GetCompanyUsecase::new(company_repo, user_repo);
        let company = usecase.get_company_usecase(get_company_req.user_id).await;
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

    pub async fn change_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(change_company_req): Json<ChangeCompanyReq>,
    ) -> Response<Body> {
        // validation for change_company_dto
        if let Err(errors) = change_company_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = ChangeCompanyUsecase::new(user_repo, company_repo);
        match usecase.change_company_usecase(change_company_req).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
                (StatusCode::OK, Json(response)).into_response()
            }
            Err(error) => {
                let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
                println!("{}", error);
                return (StatusCode::BAD_REQUEST, Json(response)).into_response();
            }
        }
    }

    pub async fn delete_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(get_company_req): Json<GetCompanyReq>,
    ) -> Response<Body> {
        // validation for get_company_dto
        if let Err(errors) = get_company_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = DeleteCompanyUsecase::new(user_repo, company_repo);
        match usecase.delete_company_usecase(get_company_req).await {
            Ok(_) => {
                let response = BasicResponse::ok("Success".to_string());
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
