extern crate blueberry_homework;

use axum::{
    routing::get, Extension, Router
};
use blueberry_homework::{config::config::Config, internal::{data::repository::{CompanyRepositoryImpl, UserRepositoryImpl}, domain::repository_interface::{company_repository::CompanyRepository, user_repository::UserRepository}}, router};
use tracing::Level;
use std::{net::SocketAddr, sync::Arc};
use tower_http::trace::{self, TraceLayer};
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    // setup logging
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::DEBUG)
        .with_target(false)
        .init();

    // Load config
    let config = Config::init_config();
    tracing::info!("Config loaded: {:?}", config);

    // 공유 레포지토리를 여기서 설정해주는 이유는
    // 현재 [] 리스트가 생성되는데 이 리스트가 앱이 종료될 때까지 유지되어야 하기 때문이다.
    // 따라서 앱이 종료될 때까지 리스트가 유지되어야 하기 때문에 여기서 설정해준다.
    let user_repo = Arc::new(UserRepositoryImpl::new()) as Arc<dyn UserRepository + Send + Sync>;
    let company_repo = Arc::new(CompanyRepositoryImpl::new()) as Arc<dyn CompanyRepository + Send + Sync>;

    // Setup base router
    let app = Router::new()
        .route("/", get(|| async { "Hello, World!" }));

    // name_router의 라우터를 가져와서 Extension 레이어 추가
    let name_router = router::name_router::create_router()
        .layer(Extension(user_repo.clone()));

    let company_router = router::company_router::create_router()
        .layer(Extension(company_repo.clone()))
        .layer(Extension(user_repo.clone()));

    // 모든 라우터를 병합한 후에 TraceLayer 적용
    let app = app
        .merge(name_router)
        .merge(company_router)
        .layer(
            TraceLayer::new_for_http()
                .make_span_with(trace::DefaultMakeSpan::new().level(Level::INFO))
                .on_response(trace::DefaultOnResponse::new().level(Level::INFO))
        );

    // Start server
    let addr = SocketAddr::from(([127, 0, 0, 1], config.server_port));
    tracing::info!("Listening on {}", addr);

    let listener = TcpListener::bind(addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
