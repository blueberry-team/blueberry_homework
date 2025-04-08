use std::sync::Mutex;
use async_trait::async_trait;

use crate::internal::domain::entities::user_entity::UserEntity;
use crate::internal::domain::repository_interface::user_repository::UserRepository;

// make inmemory list
pub struct UserRepositoryImpl {
    users: std::sync::Mutex<Vec<UserEntity>>,
}

impl UserRepositoryImpl {
    // create user repository
    pub fn new() -> Self {
        Self {
            users: Mutex::new(Vec::new()),
        }
    }
}

#[async_trait]
impl UserRepository for UserRepositoryImpl {
    fn new() -> Self {
        UserRepositoryImpl::new()
    }

    async fn create_name(&self, name: String) -> UserEntity {
        let user = UserEntity { name };
        let mut users = self.users.lock().unwrap();
        users.push(user.clone());
        user
    }

    async fn get_names(&self) -> Vec<UserEntity> {
        let names = self.users.lock().unwrap();
        names.clone()
    }

    // async fn get_names(&self) -> Vec<String> {
    //     let names = self.users.lock().unwrap();
    //     names.iter().map(|user| user.name.clone()).collect()
    // }
}
