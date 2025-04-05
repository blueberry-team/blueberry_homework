package com.example.berry_server.berry_server

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

@SpringBootApplication
class BerryServerApplication

fun main(args: Array<String>) {
    runApplication<BerryServerApplication>(*args)
}
