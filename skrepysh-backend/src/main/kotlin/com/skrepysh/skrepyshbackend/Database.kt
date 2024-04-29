package com.skrepysh.skrepyshbackend

import org.ktorm.entity.Entity
import org.ktorm.schema.Table
import org.ktorm.schema.varchar
import org.ktorm.database.Database
import org.ktorm.dsl.eq
import org.ktorm.entity.add
import org.ktorm.entity.find
import org.ktorm.entity.sequenceOf
import org.flywaydb.core.Flyway
import org.springframework.boot.runApplication



class DatabaseVM() {
    interface VirtualMachine : Entity<VirtualMachine>{
        companion object : Entity.Factory<VirtualMachine>()

        var ip: String
        var os: String
    }
    object VirtualMachinesTable : Table<VirtualMachine>("VirtualMachinesTable") {
        val ip = varchar("ip").primaryKey().bindTo { it.ip }
        val os = varchar("os").bindTo { it.os }
    }
    fun init(): Database {
        // Конфигурация базы данных
        val database = Database.connect(
            url = "jdbc:postgresql://localhost:5432/my_database",
            driver = "org.postgresql.Driver",
            user = "my_user",
            password = "my_password"
        )

        // Инициализация и запуск миграций с Flyway
        val flyway = Flyway.configure()
            .dataSource("jdbc:postgresql://localhost:5432/my_database", "my_user", "my_password")
            .load()
        flyway.migrate()

        return database
    }

    fun addVM(database: Database,ip : String, os : String) : Boolean {

         val newVM= database.sequenceOf(VirtualMachinesTable).add(
             VirtualMachine{this.ip = ip; this.os=os}
         )

        return newVM ==1
    }

    fun deleteVM(database : Database, ip : String) : Boolean {
        val VM = database.sequenceOf(VirtualMachinesTable).find { tmp -> tmp.ip eq ip }
        val affectedVMsNumber = VM?.delete()
        return affectedVMsNumber == 1
    }



}