use axum::{
    routing::{get, post, put}, Router
};

use crate::internal::handler::user_handler::UserHandler;

pub fn create_router() -> Router {
    // 공유 레포지토리 생성

    Router::new()
        .nest(
            "/user",
            Router::new()
                .route("/sign-up", post(UserHandler::sign_up_handler))
                .route("/log-in", post(UserHandler::log_in_handler))
                .route("/get-user", get(UserHandler::get_user_handler))
                .route("/change-user", put(UserHandler::change_user_handler))
        )
}
