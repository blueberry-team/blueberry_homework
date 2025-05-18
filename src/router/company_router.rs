use axum::{routing::{delete, get, post, put}, Router};

use crate::internal::handler::company_handler::CompanyHandler;

pub fn create_router() -> Router {

    Router::new()
        .route("/create-company", post(CompanyHandler::create_company_handler))
        .route("/get-company", get(CompanyHandler::get_company_handler))
        .route("/change-company", put(CompanyHandler::change_company_handler))
        .route("/delete-company", delete(CompanyHandler::delete_company_handler))
}
