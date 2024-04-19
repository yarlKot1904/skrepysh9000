package com.skrepysh.skrepyshbackend.config

import com.fasterxml.jackson.databind.ObjectMapper
import com.fasterxml.jackson.dataformat.yaml.YAMLFactory
import com.fasterxml.jackson.module.kotlin.KotlinModule
import java.io.File

fun readConfig(path: String): Config {
    val mapper = ObjectMapper(YAMLFactory())
    mapper.registerModule(KotlinModule.Builder().build())
    val yamlFile = File(path)
    return mapper.readValue(yamlFile.readBytes(), Config::class.java)
}