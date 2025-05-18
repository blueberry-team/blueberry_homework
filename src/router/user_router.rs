use axum::{
    routing::{get, post, put}, Router
};

use crate::internal::handler::user_handler::{
    sign_up_handler::SignUpHandler,
    log_in_handler::LogInHandler,
    change_user_handler::ChangeUserHandler,
    get_user_handler::GetUserHandler,
};

pub fn create_router() -> Router {
    // 공유 레포지토리 생성

    Router::new()
        .nest(
            "/user",
            Router::new()
                .route("/sign-up", post(SignUpHandler::sign_up_handler))
                .route("/log-in", post(LogInHandler::log_in_handler))
                .route("/get-user", get(GetUserHandler::get_user_handler))
                .route("/change-user", put(ChangeUserHandler::change_user_handler))
        )
}
