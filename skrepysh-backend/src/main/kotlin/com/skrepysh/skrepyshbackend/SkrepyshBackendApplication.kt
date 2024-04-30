package com.skrepysh.skrepyshbackend

import com.skrepysh.skrepyshbackend.config.Config
import com.skrepysh.skrepyshbackend.config.DatabaseConfig
import com.skrepysh.skrepyshbackend.config.readConfig
import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.context.annotation.Bean


@SpringBootApplication
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
}

fun main(args: Array<String>) {
    runApplication<SkrepyshBackendApplication>(*args)
}
