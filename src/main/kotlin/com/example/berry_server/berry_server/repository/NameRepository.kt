package com.example.berry_server.berry_server.repository

import com.example.berry_server.berry_server.dto.model.NameItem
import org.springframework.stereotype.Repository

@Repository
class NameRepository {

    private val nameList = mutableListOf<NameItem>()

    fun createName(name: String): Boolean {
        // 중복 이름 검사 | 검사 rule 추가시 반환 형태 변경
        if (nameList.any { it.name == name }) return false

        nameList.add(NameItem(name))
        return true
    }

    fun getNameList(): List<NameItem> {
        return nameList.toList()
    }


    fun deleteName(index: Int): Boolean {
        // 인덱스 범위 검사
        if (index !in nameList.indices) return false

        nameList.removeAt(index)
        return true
    }
}