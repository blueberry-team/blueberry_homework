use axum::{
    middleware, routing::{delete, get, post, put}, Router
};

use crate::internal::{
    handler::company_handler::{
        change_company_handler::ChangeCompanyHandler,
        create_company_handler::CreateCompanyHandler,
        delete_company_handler::DeleteCompanyHandler,
        get_company_handler::GetCompanyHandler
    },
    middlewares::verify_token_middleware::verify_token_middleware
};

pub fn create_router(jwt_secret_key: String) -> Router {

    // create private router (needs token)
    let private_router = Router::new()
        .route("/create-company", post(CreateCompanyHandler::create_company_handler))
        .route("/get-company", get(GetCompanyHandler::get_company_handler))
        .route("/change-company", put(ChangeCompanyHandler::change_company_handler))
        .route("/delete-company", delete(DeleteCompanyHandler::delete_company_handler))
        .layer(middleware::from_fn_with_state(
            jwt_secret_key.clone(),
            verify_token_middleware
        ));

    // merge public and private router
    Router::new()
        .nest("/company", private_router)
}
