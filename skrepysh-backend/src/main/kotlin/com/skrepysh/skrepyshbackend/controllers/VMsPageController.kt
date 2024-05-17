package com.skrepysh.skrepyshbackend.controllers

import com.fasterxml.jackson.core.JsonProcessingException
import com.fasterxml.jackson.databind.ObjectMapper
import com.skrepysh.skrepyshbackend.database.DatabaseVM
import jakarta.servlet.http.HttpServletRequest
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import org.springframework.beans.factory.annotation.Autowired
import org.springframework.stereotype.Controller
import org.springframework.ui.Model
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PathVariable
import org.springframework.web.bind.annotation.RequestParam
import org.springframework.web.client.RestTemplate


@Controller
class VMsPageController(@Autowired private val database: DatabaseVM) {
    @Autowired
    private val context: HttpServletRequest? = null
    val log: Logger = LoggerFactory.getLogger(VMsPageController::class.java)

    @GetMapping("/vms")
    fun vmsPage(
        @RequestParam(name = "page", defaultValue = "1") page: Int,
        model: Model
    ): String {
        log.info("${context!!.method} request /vms")

        val limit = 10
        val offset = (page - 1) * limit
        val vms = database.listVMs(offset, limit)
        val totalVMs = database.getVMsCount()

        model.addAttribute("vms", vms)
        model.addAttribute("currentPage", page)
        model.addAttribute("totalPages", (totalVMs + limit - 1) / limit)

        return "vms"
    }

    @GetMapping("/vms/{id}")
    fun getVirtualMachineById(@PathVariable id: Int, model: Model): String {
        log.info("${context!!.method} request /vms/$id")
        val vm = database.getVMByID(id)
        if (vm != null) {
            model.addAttribute("vm", vm)

            val metrics = fetchMetrics(vm.ip!!)
            if (metrics != null) {
                model.addAttribute("metricsJson", metrics)
            } else {
                model.addAttribute("metricsJson", "Metrics are not available")
            }

            return "vm-detail"
        } else {
            log.error("vm with id $id not found")
            return "vm-not-found"
        }
    }

    private fun fetchMetrics(ip: String): String? {
        val restTemplate = RestTemplate()
        val url = "http://$ip:48934/metrics"
        return try {
            val objectMapper = ObjectMapper()
            objectMapper.writeValueAsString(
                restTemplate.getForObject(url, Map::class.java) as Map<String, Any>?
            )
        } catch (e: Exception) {
            null
        }
    }
}
