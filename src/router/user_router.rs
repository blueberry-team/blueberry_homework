use axum::{
    middleware,
    routing::{get, post, put},
    Router
};

use crate::internal::{
    handler::user_handler::{
        change_user_handler::ChangeUserHandler,
        get_user_handler::GetUserHandler,
        log_in_handler::LogInHandler,
        refresh_token_handler::RefreshTokenHandler,
        sign_up_handler::SignUpHandler
    },
    middlewares::verify_token_middleware::verify_token_middleware
};

pub fn create_router(jwt_secret_key: String) -> Router {

    // create public router (does not need token)
    let public_router = Router::new()
        .route("/sign-up", post(SignUpHandler::sign_up_handler))
        .route("/log-in", post(LogInHandler::log_in_handler))
        .with_state(jwt_secret_key.clone());

    // create private router (needs token)
    let private_router = Router::new()
        .route("/get-user", get(GetUserHandler::get_user_handler))
        .route("/change-user", put(ChangeUserHandler::change_user_handler))
        .route("/refresh-token", post(RefreshTokenHandler::refresh_token_handler))
        .with_state(jwt_secret_key.clone())
        .layer(middleware::from_fn_with_state(
            jwt_secret_key.clone(),
            verify_token_middleware
        ));

    // merge public and private router
    Router::new()
        .nest("/user", public_router.merge(private_router))
}
