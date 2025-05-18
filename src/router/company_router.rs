use axum::{
    routing::{delete, get, post, put},
    Router,
};

use crate::internal::handler::company_handler::{
    create_company_handler::CreateCompanyHandler,
    get_company_handler::GetCompanyHandler,
    change_company_handler::ChangeCompanyHandler,
    delete_company_handler::DeleteCompanyHandler,
};

pub fn create_router() -> Router {

    Router::new()
        .nest(
            "/company",
            Router::new()
                .route("/create-company", post(CreateCompanyHandler::create_company_handler))
                .route("/get-company", get(GetCompanyHandler::get_company_handler))
                .route("/change-company", put(ChangeCompanyHandler::change_company_handler))
                .route("/delete-company", delete(DeleteCompanyHandler::delete_company_handler))
        )
}
