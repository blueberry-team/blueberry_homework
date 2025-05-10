use std::{error::Error, sync::Arc};

use scylla::client::session::Session;

use crate::{config::config::Config, db::scylla_db::ScyllaDB};

pub struct App {
    pub config: Config,
    // it means that the scylla db is not initialized yet
    pub scylla_db: Option<ScyllaDB>,
    pub session: Arc<Session>,
}

impl App {
    pub async fn new() -> Result<Self, Box<dyn Error>> {
        let config = Config::init_config();

        let scylla_db = ScyllaDB::new(
            config.scylla_db_port,
            config.scylla_db_host.clone(),
            config.scylla_db_keyspace.clone(),
            config.scylla_db_user.clone(),
            config.scylla_db_password.clone(),
        );

        let session = scylla_db.init_scylla_db().await.unwrap();

        Ok(Self {
            config: config,
            scylla_db: Some(scylla_db),
            session: Arc::new(session),
        })
    }

    // create a new app with the given components
    pub fn with_components(
        config: Config,
        scylla_db: ScyllaDB,
        session: Session,
    ) -> Self {
        Self {
            config,
            scylla_db: Some(scylla_db),
            session: Arc::new(session),
        }
    }

    // get session helper method
    pub fn session(&self) -> &Session {
        &self.session
    }
}
