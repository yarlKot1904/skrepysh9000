package com.skrepysh.skrepyshbackend.database

import com.skrepysh.skrepyshbackend.config.DatabaseConfig
import org.ktorm.database.Database
import org.ktorm.dsl.eq
import org.ktorm.dsl.delete
import org.ktorm.dsl.insert
import org.springframework.beans.factory.annotation.Autowired
import java.sql.SQLException


class DatabaseVM(@Autowired private val dbConf: DatabaseConfig) {
    private var database: Database

    init {
        val url = "jdbc:postgresql://${dbConf.host}:${dbConf.port}/${dbConf.databaseName}"
        val dbPassword = System.getenv(dbConf.passwordEnv) ?: throw RuntimeException("database password not set")

        this.database = Database.connect(
            url = url,
            driver = "org.postgresql.Driver",
            user = dbConf.user,
            password = dbPassword,
        )
    }

    fun addVM(ip: String, os: String) {
        database.insert(VirtualMachinesTable) {
            set(it.ip, ip)
            set(it.os, os)
            set(it.isActive, true)
        }
    }

    fun deleteVM(ip: String) {
        database.delete(VirtualMachinesTable) {
            it.ip eq ip
        }
    }


}