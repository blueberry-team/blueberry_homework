package com.example.berry_server.berry_server.controller

import com.example.berry_server.berry_server.repository.NameRepository
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
    fun createName(@RequestBody name: String): String {
        nameRepository.createName(name)
        return "이름 등록: $name"
    }

    // 이름 전체 조회
    @GetMapping("/getName")
    fun getNames(): List<String> {
        return nameRepository.getNameList()
    }
}