package com.skrepysh.skrepyshbackend.database

import org.ktorm.entity.Entity
import org.ktorm.schema.Table
import org.ktorm.schema.boolean
import org.ktorm.schema.varchar
import org.ktorm.schema.int

interface VirtualMachine : Entity<VirtualMachine> {
    companion object : Entity.Factory<VirtualMachine>()

    var id: Int
    var ip: String
    var os: String
    var isActive: Boolean
}

object VirtualMachinesTable : Table<VirtualMachine>("virtual_machines") {
    var id = int("id").primaryKey().bindTo { it.id }
    var ip = varchar("ip").bindTo { it.ip }
    var os = varchar("os").bindTo { it.os }
    var isActive = boolean("is_active").bindTo { it.isActive }
}