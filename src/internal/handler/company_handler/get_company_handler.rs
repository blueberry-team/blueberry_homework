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
        req::company_req::GetCompanyReq,
        res::basic_response::BasicResponse,
    },
    internal::domain::{
        repository_interface::{
            company_repository::CompanyRepository,
            user_repository::UserRepository
        },
        usecases::company_usecases::get_company_usecase::GetCompanyUsecase,
    },
};

pub struct GetCompanyHandler;

impl GetCompanyHandler {

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

}
