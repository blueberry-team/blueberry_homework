use serde::Deserialize;
use dotenv::dotenv;
use std::env;

#[derive(Debug, Deserialize)]
pub struct Config {
    pub server_port: u16,
    pub scylla_db_port: u16,
    pub scylla_db_host: String,
    pub scylla_db_user: String,
    pub scylla_db_password: String,
    pub scylla_db_keyspace: String,
}

impl Config {
    pub fn init_config() -> Self {
        dotenv().ok();

        let server_port = env::var("SERVER_PORT")
            .unwrap_or_else(|_| "8090".to_string())
            .parse::<u16>()
            .expect("SERVER_PORT must be a number");

        let scylla_db_port = env::var("SCYLLA_DB_PORT")
            .unwrap_or_else(|_| "9042".to_string())
            .parse::<u16>()
            .expect("SCYLLA_DB_PORT must be a number");

        let scylla_db_host = env::var("SCYLLA_DB_HOST")
            .unwrap_or_else(|_| "localhost".to_string());

        let scylla_db_user = env::var("SCYLLA_DB_USER")
            .unwrap_or_else(|_| "cassandra".to_string());

        let scylla_db_password = env::var("SCYLLA_DB_PASSWORD")
            .unwrap_or_else(|_| "cassandra".to_string());

        let scylla_db_keyspace = env::var("SCYLLA_DB_KEYSPACE")
            .unwrap_or_else(|_| "test".to_string());

        Config {
            server_port,
            scylla_db_port,
            scylla_db_host,
            scylla_db_user,
            scylla_db_password,
            scylla_db_keyspace,
        }
    }
}
