use std::sync::Mutex;
use async_trait::async_trait;
use chrono::{DateTime, Utc};
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

    async fn create_name(&self, user: UserEntity) -> Result<UserEntity, String> {
        let mut users = self.users.lock().unwrap();
        // check name if user already have same value
        if users.iter().any(|u| u.name == user.name) {
            return Err(format!("A name with the same value already exists"));
        }
        users.push(user.clone());
        Ok(user)
    }

    async fn change_name(&self, user_id: String, name: String, update_at: DateTime<Utc>) -> Result<(), String> {
        let mut users = self.users.lock().unwrap();
        // search name
        let position = users.iter().position(|user| user.id == user_id);
        match position {
            Some(index) => {
                // check name if user already have same value
                if users[index].name == name {
                    return Err(format!("The name is the same as the current name"));
                }
                // if another user have same name
                if users.iter().any(|user| user.name == name) {
                    return Err(format!("A name with the same value already exists"));
                }
                users[index].name = name;
                users[index].update_at = update_at;
                Ok(())
            },
            None => Err(format!("User not found: {}", user_id))
        }
    }

    async fn find_by_name(&self, name: String) -> Result<(), String> {
        let users = self.users.lock().unwrap();
        let position = users.iter().position(|user| user.name == name);
        match position {
            Some(_) => Ok(()),
            None => Err(format!("User not found: {}", name))
        }
    }

    async fn get_names(&self) -> Vec<UserEntity> {
        let names = self.users.lock().unwrap();
        names.clone()
    }

    // async fn get_names(&self) -> Vec<String> {
    //     let names = self.users.lock().unwrap();
    //     names.iter().map(|user| user.name.clone()).collect()
    // }
    async fn delete_index(&self, index: u32) -> Result<(), String> {
        let mut users = self.users.lock().unwrap();
        let index = index as usize;

        if index >= users.len() {
            return Err(format!("Index out of bounds: {}", index));
        }

        users.remove(index);
        Ok(())
    }

    async fn delete_name(&self, name: String) -> Result<(), String> {
        let mut users = self.users.lock().unwrap();

        // search index
        let position = users.iter().position(|user| user.name == name);

        // if not found return err or delete
        match position {
            Some(index) => {
                users.remove(index);
                Ok(())
            },
            None => Err(format!("User not found: {}", name))
        }
    }
}
