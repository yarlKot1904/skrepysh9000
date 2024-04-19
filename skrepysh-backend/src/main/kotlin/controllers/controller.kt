package controllers

import org.springframework.stereotype.Controller
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody

data class IpAddress(var ip: String)

@Controller
class HtmlController {

    @PostMapping("/init")
    fun getIP(@RequestBody request : IpAddress): String {
        val ip = request.ip
        return ip
    }
}