package com.example.berry_server.util

// Validation Class
object Validation {

    fun validateName(name: String): String? {
        if (name.isEmpty()) return Messages.ERROR_CREATE_NAME_EMPTY
        if (name.length > 50) return Messages.ERROR_CREATE_NAME_UNDER_50

        return null
    }
}