package com.skrepysh.skrepyshbackend.controllers

import com.skrepysh.skrepyshbackend.database.DatabaseVM
import jakarta.servlet.http.HttpServletRequest
import org.slf4j.Logger
import org.slf4j.LoggerFactory
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
    @Autowired
    private val context: HttpServletRequest? = null

    val log: Logger = LoggerFactory.getLogger(InitController::class.java)

    @PostMapping("/init")
    @ResponseBody
    fun init(@RequestBody request: InitRequestBody): ResponseEntity<String> {
        log.info("${context!!.method} request /init: $request")
        try {
            database.addVM(request.ip, request.os)
            return ResponseEntity<String>(HttpStatus.OK)
        } catch (e: Exception) {
            log.error("Error adding vm to database: $request")
            throw ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "error: ${e.message}")
        }

    }
}