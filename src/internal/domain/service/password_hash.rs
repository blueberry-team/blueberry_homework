use argon2::{password_hash::{
    rand_core::OsRng,
    SaltString,
    },
    Argon2
};


pub fn hash_password(password: &str) -> Result<(Vec<u8>, String), String> {
    let pssword_bytes = password.as_bytes();
    let salt = SaltString::generate(&mut OsRng);
    let salt_str = salt.as_str().to_string();
    let mut output_key_material = [0u8; 32];

    let argon2 = Argon2::default();

    argon2.hash_password_into(
        pssword_bytes,
        salt_str.as_bytes(),
        &mut output_key_material,
    )
    .map_err(|e| format!("Error hashing password: {}", e))?;


    Ok((output_key_material.to_vec(), salt_str))
}

pub fn verify_password(password: &str, hash: &[u8], salt: &str) -> Result<bool, String> {
    // hash is hashed password with salt
    let mut output_key_material = [0u8; 32];

    // input password and salt make new hash
        Argon2::default()
            .hash_password_into(
                password.as_bytes(),
                salt.as_bytes(),
                &mut output_key_material,
            )
            .map_err(|e| format!("Error hashing password: {}", e))?;

    // compare new hash and hashed password with salt
    Ok(output_key_material.to_vec() == hash)
}
