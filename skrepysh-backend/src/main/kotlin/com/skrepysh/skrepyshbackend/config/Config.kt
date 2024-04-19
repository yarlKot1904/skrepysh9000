package com.skrepysh.skrepyshbackend.config

import com.fasterxml.jackson.databind.PropertyNamingStrategies
import com.fasterxml.jackson.databind.annotation.JsonNaming

@JsonNaming(PropertyNamingStrategies.KebabCaseStrategy::class)
data class Config(
    val database: Database
)

@JsonNaming(PropertyNamingStrategies.KebabCaseStrategy::class)
data class Database(
    val host: String,
    val port: Short,
    val user: String,
    val passwordEnv: String,
)