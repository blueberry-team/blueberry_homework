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
        req::company_req::DeleteCompanyReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        },
        usecases::company_usecases::delete_company_usecase::DeleteCompanyUsecase,
    },
};

pub struct DeleteCompanyHandler;

impl DeleteCompanyHandler {

    pub async fn delete_company_handler(
        Extension(company_repo): Extension<Arc<dyn CompanyRepository + Send + Sync>>,
        Extension(user_repo): Extension<Arc<dyn UserRepository + Send + Sync>>,
        Json(delete_company_req): Json<DeleteCompanyReq>,
    ) -> Response<Body> {
        // validation for get_company_dto
        if let Err(errors) = delete_company_req.validate() {
            let response = BasicResponse::bad_request(format!("error"), format!("name must be 1 and 50 characters"));
            println!("{}", errors);
            return (StatusCode::BAD_REQUEST, Json(response)).into_response();
        }
        let usecase = DeleteCompanyUsecase::new(user_repo, company_repo);
        match usecase.delete_company_usecase(delete_company_req).await {
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
