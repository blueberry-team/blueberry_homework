package com.example.berry_server.dto.response

sealed class BaseApiResponse

data class ApiResponse<T>(
    val message: String,
    val data: T? = null
): BaseApiResponse()

data class ApiErrorResponse(
    val error: String,
    val message: String
) : BaseApiResponse()