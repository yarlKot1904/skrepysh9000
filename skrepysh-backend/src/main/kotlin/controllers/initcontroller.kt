package controllers

import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RestController

data class IpAddress(var ip: String)

@RestController
class InitController {

    @PostMapping("/init")
    fun init(@RequestBody request : IpAddress): String {
        val ip = request.ip
        return ip
    }
}