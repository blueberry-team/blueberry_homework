use axum::{routing::post, Router};

use crate::handler::company_handler::CompanyHandler;

pub fn create_router() -> Router {

    Router::new()
        .route("/create-company", post(CompanyHandler::create_company_handler))
}
