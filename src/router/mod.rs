pub mod user_router;
pub mod company_router;

use axum::{Extension, Router};
pub use user_router::create_router as name_router;
pub use company_router::create_router as company_router;
use tower_http::trace::{self, TraceLayer};
use tracing::Level;

use crate::di::AppDI;

pub fn create_app_router(app_state: AppDI) -> Router {
    Router::new()
        .merge(user_router::create_router())
        .merge(company_router::create_router())
        .layer(Extension(app_state.user_repo.clone()))
        .layer(Extension(app_state.company_repo.clone()))
        .layer(
            TraceLayer::new_for_http()
                .make_span_with(trace::DefaultMakeSpan::new().level(Level::INFO))
                .on_response(trace::DefaultOnResponse::new().level(Level::INFO))
        )
}
