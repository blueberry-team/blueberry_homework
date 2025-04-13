package com.example.berry_server.controller

import com.example.berry_server.util.Constants
import com.example.berry_server.dto.request.name.CreateNameRequest
import com.example.berry_server.dto.request.name.DeleteNameRequest
import com.example.berry_server.dto.response.ApiErrorResponse
import com.example.berry_server.dto.response.ApiResponse
import com.example.berry_server.dto.response.BaseApiResponse
import com.example.berry_server.repository.NameRepository
import com.example.berry_server.util.Messages
import com.example.berry_server.util.Validation
import org.springframework.web.bind.annotation.DeleteMapping
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping("/name")
class NameController(
    private val nameRepository: NameRepository
) {

    // 이름 생성
    @PostMapping("/createName")
    fun createName(@RequestBody request: CreateNameRequest): BaseApiResponse {
        // null 체크
        if (request.name == null) {
            return ApiErrorResponse(
                message = Constants.ERROR,
                error = Messages.ERROR_CREATE_NAME_INPUT_NAME
            )
        }

        // 이름 validation
        Validation.validateName(request.name)?.let { errorMessage ->
            return ApiErrorResponse(
                message = Constants.ERROR,
                error = errorMessage
            )
        }

        // 중복 검사
        if (nameRepository.existName(request.name)) {
            return ApiErrorResponse(
                message = Constants.ERROR,
                error = Messages.ERROR_CREATE_NAME_DUPLICATED
            )
        }

        // 이름 생성 성공
        nameRepository.createName(request.name)
        return ApiResponse(
            message = "${Messages.SUCCESS_CREATE_NAME}: ${request.name}",
            data = nameRepository.getNameList()
        )
    }

    // 이름 전체 조회
    @GetMapping("/getName")
    fun getName(): BaseApiResponse {
        return ApiResponse(
            message = Constants.SUCCESS,
            data = nameRepository.getNameList()
        )
    }

    // 인덱스 이름 삭제
    @DeleteMapping("/deleteName")
    fun deleteName(@RequestBody request: DeleteNameRequest): BaseApiResponse {
        if (request.deleteIndex == null) {
            return ApiErrorResponse(
                message = Constants.ERROR,
                error = Messages.ERROR_DELETE_NAME_INDEX_NULL,
            )
        }

        val isDeleted = nameRepository.deleteName(request.deleteIndex)

        if (!isDeleted) {
            return ApiErrorResponse(
                message = Constants.ERROR,
                error = "${Messages.ERROR_DELETE_NAME_INVALID_INDEX}: ${request.deleteIndex}",
            )
        }

        return ApiResponse(
            message = Constants.SUCCESS,
            data = nameRepository.getNameList()
        )
    }
}