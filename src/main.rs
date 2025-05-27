extern crate blueberry_homework;

use blueberry_homework::{
    app::app::App,
    di::AppDI,
    router
};
use std::net::SocketAddr;
use tokio::net::TcpListener;

#[tokio::main]
async fn main() {
    // setup logging
    tracing_subscriber::fmt()
        .with_max_level(tracing::Level::DEBUG)
        .with_target(false)
        .init();

    // Load config
    let app = App::new().await.unwrap();
    tracing::info!("App loaded: {}", app.config.server_port.clone());

    // initialize app state dependency injection
    let app_state = AppDI::new(app.session.clone(), app.config.jwt_secret_key.clone());

    // initialize app router
    let app_router = router::create_app_router(app_state);

    // Start server
    let addr = SocketAddr::from(([127, 0, 0, 1], app.config.server_port));
    tracing::info!("Listening on {}", addr);

    let listener = TcpListener::bind(addr).await.unwrap();
    axum::serve(listener, app_router).await.unwrap();
}
