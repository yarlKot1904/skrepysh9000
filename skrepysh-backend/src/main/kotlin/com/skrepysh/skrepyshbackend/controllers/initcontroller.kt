package com.skrepysh.skrepyshbackend.controllers

import com.skrepysh.skrepyshbackend.database.DatabaseVM
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.ResponseBody
import org.springframework.web.bind.annotation.RestController
import org.springframework.web.server.ResponseStatusException


data class InitRequestBody(var ip: String, var os: String)

@RestController
class InitController(@Autowired private val database: DatabaseVM) {
    @PostMapping("/init")
    @ResponseBody
    fun init(@RequestBody request: InitRequestBody): ResponseEntity<String> {
        try {
            database.addVM(request.ip, request.os)
            return ResponseEntity<String>(request.ip, HttpStatus.OK)
        } catch (e: Exception) {
            throw ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "error: ${e.message}")
        }

    }


}