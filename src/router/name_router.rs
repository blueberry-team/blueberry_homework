
use axum::{
    routing::{delete, get, post, put}, Router
};

use crate::internal::handler::user_handler::UserHandler;

pub fn create_router() -> Router {
    // 공유 레포지토리 생성

    Router::new()
        .route("/create-name", post(UserHandler::create_user_handler))
        .route("/change-name", put(UserHandler::change_name_handler))
        .route("/get-names", get(UserHandler::get_names_handler))
        .route("/delete-name", delete(UserHandler::delete_name_handler))
}
