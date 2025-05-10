// module for dependency injection when creating a new app

use std::sync::Arc;

use scylla::client::session::Session;

use crate::internal::{
    data::repository::{
        CompanyRepositoryImpl,
        UserRepositoryImpl
    },
    domain::repository_interface::{
        company_repository::CompanyRepository,
        user_repository::UserRepository
    }
};

// dependency injection for the app
pub struct AppDI {
    // user repository
    pub user_repo: Arc<dyn UserRepository + Send + Sync>,

    // company repository
    pub company_repo: Arc<dyn CompanyRepository + Send + Sync>,
}

impl AppDI {
    pub fn new(session: Arc<Session>) -> Self {
        Self {
            // user repository
            user_repo: Arc::new(UserRepositoryImpl::new(session.clone())) as Arc<dyn UserRepository + Send + Sync>,

            // company repository
            company_repo: Arc::new(CompanyRepositoryImpl::new(session.clone())) as Arc<dyn CompanyRepository + Send + Sync>,
        }
    }
}
