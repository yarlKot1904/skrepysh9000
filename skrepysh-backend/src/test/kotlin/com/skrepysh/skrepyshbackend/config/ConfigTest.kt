package com.skrepysh.skrepyshbackend.config

import org.junit.jupiter.api.Test

import org.junit.jupiter.api.Assertions.*

class ConfigTest {
    val configPath = "src/test/kotlin/com/skrepysh/skrepyshbackend/config/files/config.yaml"

    @Test
    fun readConfig() {
        val config = readConfig(configPath)
        assertEquals("localhost", config.database.host)
        assertEquals(5432, config.database.port)
        assertEquals("db-name", config.database.databaseName)
        assertEquals("pg-user", config.database.user)
        assertEquals("DB_PASSWORD", config.database.passwordEnv)
    }
}