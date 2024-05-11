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


data class DeleteRequestBody(var ip: String)

@RestController
class DeleteController(@Autowired private val database: DatabaseVM) {
    @Autowired
    private val context: HttpServletRequest? = null

    val log: Logger = LoggerFactory.getLogger(DeleteController::class.java)

    @PostMapping("/delete")
    @ResponseBody
    fun delete(@RequestBody request: DeleteRequestBody): ResponseEntity<String> {
        log.info("${context!!.method} request /delete: $request")
        try {
            database.deleteVM(request.ip)
            return ResponseEntity<String>(HttpStatus.OK)
        } catch (e: Exception) {
            log.error("Error deleting vm from database: $request")
            throw ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "error: ${e.message}")
        }

    }
}