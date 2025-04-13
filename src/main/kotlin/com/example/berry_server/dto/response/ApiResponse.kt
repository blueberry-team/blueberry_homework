package com.example.berry_server.dto.response

data class ApiResponse<T>(
    val message: String,
    val data: T? = null,
    val error: String? = null
)