package com.example.berry_server.berry_server.repository

import org.springframework.stereotype.Repository

@Repository
class NameRepository {

    private val nameList = mutableListOf<String>()

    fun createName(name: String) {
        nameList.add(name)
    }

    fun getNameList(): List<String> {
        return nameList.toList()
    }
}