package com.skrepysh.skrepyshbackend.database

import com.skrepysh.skrepyshbackend.config.DatabaseConfig
import org.ktorm.database.Database
import org.ktorm.dsl.*
import org.ktorm.entity.sequenceOf
import org.ktorm.entity.toList
import org.springframework.beans.factory.annotation.Autowired
import java.util.*


class DatabaseVM(@Autowired private val dbConf: DatabaseConfig) {
    private var database: Database

    data class VirtualMachineEntity(var ip: String?, var os: String?)

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


    fun toList(): List<VirtualMachineEntity> {
        val query = database.from(VirtualMachinesTable)
            .select(
                VirtualMachinesTable.ip,
                VirtualMachinesTable.os,
            )
            .where { VirtualMachinesTable.isActive eq true }
            .map { row ->
                VirtualMachineEntity(row[VirtualMachinesTable.ip], row[VirtualMachinesTable.os])
            }
        return query
    }


}