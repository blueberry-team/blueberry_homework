use std::{error::Error, time::Duration};

use scylla::client::{
    session::Session,
    session_builder::SessionBuilder,
};
use serde::Deserialize;

#[derive(Debug, Deserialize)]
pub struct ScyllaDB{
    pub db_port: u16,
    pub host: String,
    pub keyspace: String,
    pub db_user: String,
    pub db_password: String,
}

impl ScyllaDB {
    // create a new scylla db instance
    pub fn new(db_port: u16, host: String, keyspace: String, db_user: String, db_password: String) -> Self {
        Self {
            db_port,
            host,
            keyspace,
            db_user,
            db_password,
        }
    }

    // database connection
    pub async fn init_scylla_db(&self) -> Result<Session, Box<dyn Error>> {
        // create a address connection string
        let node = format!("{}:{}", self.host, self.db_port);

        // connect to the scylla db
        let session: Session = SessionBuilder::new()
            .known_node(node)
            .connection_timeout(Duration::from_secs(50))
            .user(self.db_user.clone(), self.db_password.clone())
            .build()
            .await?;

        // create a keyspace
        session
            .query_iter(
                format!(
                    "CREATE KEYSPACE IF NOT EXISTS {}
                    WITH REPLICATION = {{'class': 'SimpleStrategy', 'replication_factor': 1}}",
                    self.keyspace
                ),
                &[],
            )
            .await?;

        // use the keyspace
        session.use_keyspace(&self.keyspace, false).await?;

        // db 초기화 로직, 테스트용임
        // session
        // .query_iter("DROP TABLE IF EXISTS user", &[])
        // .await?;

        session
        .query_iter("DROP TABLE IF EXISTS company", &[])
        .await?;

        // create table
        self.create_user_table(&session).await?;
        self.create_company_table(&session).await?;

        println!("Connected to ScyllaDB");

        // return the session
        Ok(session)
    }

    pub async fn create_user_table(&self, session: &Session) -> Result<(), Box<dyn Error>> {
        // create user table
        session.query_iter(
            "CREATE TABLE IF NOT EXISTS user (
            id UUID,
            name TEXT,
            email TEXT,
            password BLOB,
            salt TEXT,
            role TEXT,
            created_at BIGINT,
            updated_at BIGINT,
            PRIMARY KEY (id)
            )",
            &[],
        )
        .await?;

        // name as the pirmary key, but the problem is that a primary key cannot be changed
        // so cant update the name
        // will create a secondary index on the name column
        session.query_iter(
            "CREATE INDEX IF NOT EXISTS email_index ON user (email)",
            &[],
        )
        .await?;

        Ok(())
    }

    pub async fn create_company_table(&self, session: &Session) -> Result<(), Box<dyn Error>> {
        // create company table
        session.query_iter(
            "CREATE TABLE IF NOT EXISTS company (
            id UUID PRIMARY KEY,
            user_id UUID,
            company_name TEXT,
            company_address TEXT,
            total_staff SMALLINT,
            created_at BIGINT,
            updated_at BIGINT
            )",
            &[],
        )
        .await?;

        session.query_iter(
            "CREATE INDEX IF NOT EXISTS user_id_index ON company (user_id)",
            &[],
        )
        .await?;

        Ok(())
    }
}
