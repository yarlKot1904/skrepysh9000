package com.skrepysh.skrepyshbackend.controllers

import com.skrepysh.skrepyshbackend.database.DatabaseVM
import com.skrepysh.skrepyshbackend.database.VirtualMachine
import jakarta.servlet.http.HttpServletRequest
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.http.HttpStatus
import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.ResponseBody
import org.springframework.web.bind.annotation.RestController
import org.springframework.web.server.ResponseStatusException


@RestController
class ListVMController(@Autowired private val database: DatabaseVM) {
    @Autowired
    private val context: HttpServletRequest? = null

    val log: Logger = LoggerFactory.getLogger(ListVMController::class.java)

    @GetMapping("/listVMs")
    @ResponseBody
    fun list(): ResponseEntity<List<VirtualMachine>> {
        log.info("${context!!.method} request /listVMs")
        try {
            return ResponseEntity(database.toList(), HttpStatus.OK)
        } catch (e: Exception) {
            log.error("Error listing vm in database")
            throw ResponseStatusException(HttpStatus.INTERNAL_SERVER_ERROR, "error: ${e.message}")
        }
    }

}
