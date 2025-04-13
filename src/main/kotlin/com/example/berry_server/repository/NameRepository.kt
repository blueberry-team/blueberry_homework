package com.example.berry_server.repository

import com.example.berry_server.dto.model.NameItem
import org.springframework.stereotype.Repository

@Repository
class NameRepository {

    private val nameList = mutableListOf<NameItem>()

    fun createName(name: String) {
        nameList.add(NameItem(name))
    }

    // 중복 이름 검사
    fun existName(name: String): Boolean {
        return !nameList.any { it.name == name }
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