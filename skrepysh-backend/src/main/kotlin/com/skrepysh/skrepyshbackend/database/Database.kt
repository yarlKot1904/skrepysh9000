package com.skrepysh.skrepyshbackend.database

import com.skrepysh.skrepyshbackend.config.DatabaseConfig
import org.ktorm.database.Database
import org.ktorm.dsl.eq
import org.ktorm.entity.add
import org.ktorm.entity.find
import org.ktorm.entity.sequenceOf
import org.flywaydb.core.Flyway


class DatabaseVM(private val dbConf: DatabaseConfig) {
    private var database: Database

    init {
        val url = "jdbc:postgresql://${dbConf.host}:${dbConf.port}/${dbConf.databaseName}"
        val dbPassword = System.getenv(dbConf.passwordEnv) ?: throw RuntimeException("database password not set")

        // Конфигурация базы данных
        this.database = Database.connect(
            url = url,
            driver = "org.postgresql.Driver",
            user = dbConf.user,
            password = dbPassword,
        )

        // Инициализация и запуск миграций с Flyway
        val flyway = Flyway.configure()
            .dataSource(url, dbConf.user, dbPassword)
            .load()
        flyway.migrate()
    }

    fun addVM(ip: String, os: String): Boolean {
        val newVM = this.database.sequenceOf(VirtualMachinesTable).add(
            VirtualMachine { this.ip = ip; this.os = os }
        )
        return newVM == 1
    }

    fun deleteVM(database: Database, ip: String): Boolean {
        val VM = database.sequenceOf(VirtualMachinesTable).find { tmp -> VirtualMachinesTable.ip eq ip }
        val affectedVMsNumber = VM?.delete()
        return affectedVMsNumber == 1
    }


}