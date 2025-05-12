use argon2::{password_hash::{
    rand_core::OsRng,
    SaltString,
    }, Argon2
};


pub fn hash_password(password: &str) -> Result<Vec<u8>, String> {
    let pssword_bytes = password.as_bytes();
    let salt = SaltString::generate(&mut OsRng);

    let mut output_key_material = [0u8; 32];

    Argon2::default()
        .hash_password_into(
            pssword_bytes,
            salt.as_str().as_bytes(),
            &mut output_key_material,
        )
        .map_err(|e| format!("Error hashing password: {}", e))?;

    Ok(output_key_material.to_vec())
}
