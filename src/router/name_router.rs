
use axum::{
    routing::{get, post}, Router
};

use crate::handler::user_handler::UserHandler;

pub fn create_router() -> Router {
    // 공유 레포지토리 생성

    Router::new()
        .route("/create-name", post(UserHandler::create_user_handler))
        .route("/get-names", get(UserHandler::get_names_handler))
}
