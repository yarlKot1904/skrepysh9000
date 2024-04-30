package com.skrepysh.skrepyshbackend.config

import com.fasterxml.jackson.databind.PropertyNamingStrategies
import com.fasterxml.jackson.databind.annotation.JsonNaming

@JsonNaming(PropertyNamingStrategies.KebabCaseStrategy::class)
data class Config(
    val database: DatabaseConfig
)

@JsonNaming(PropertyNamingStrategies.KebabCaseStrategy::class)
data class DatabaseConfig(
    val host: String,
    val port: Short,
    val databaseName: String,
    val user: String,
    val passwordEnv: String,
)