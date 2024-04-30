package com.skrepysh.skrepyshbackend.database

import org.ktorm.entity.Entity
import org.ktorm.schema.Table
import org.ktorm.schema.varchar

interface VirtualMachine : Entity<VirtualMachine> {
    companion object : Entity.Factory<VirtualMachine>()

    var ip: String
    var os: String
}

object VirtualMachinesTable : Table<VirtualMachine>("VirtualMachines") {
    val ip = varchar("ip").primaryKey().bindTo { it.ip }
    val os = varchar("os").bindTo { it.os }
}