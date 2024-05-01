package com.skrepysh.skrepyshbackend

import ch.qos.logback.classic.Level
import com.skrepysh.skrepyshbackend.config.Config
import com.skrepysh.skrepyshbackend.config.DatabaseConfig
import com.skrepysh.skrepyshbackend.config.readConfig
import com.skrepysh.skrepyshbackend.database.DatabaseVM
import org.slf4j.LoggerFactory
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean


@SpringBootApplication(scanBasePackages = ["com.skrepysh.skrepyshbackend.controllers"])
class SkrepyshBackendApplication {
    @Bean
    fun config(): Config {
        val configPath = System.getenv("CONFIG_PATH") ?: throw RuntimeException("CONFIG_PATH env var not set")
        val config = readConfig(configPath)
        return config
    }

    @Bean
    fun databaseConfig(): DatabaseConfig {
         return config().database
    }

    @Bean
    fun database(): DatabaseVM {
        return DatabaseVM(databaseConfig())
    }
}

fun main(args: Array<String>) {
    //чтобы логгер правильно определялся
    val configPath = SkrepyshBackendApplication::class.java.getResource("/logback.xml")?.toString()
    System.setProperty("logging.config", configPath)

    runApplication<SkrepyshBackendApplication>(*args)
}
