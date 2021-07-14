package com.growth

import io.ktor.server.engine.*
import io.ktor.server.netty.*
import com.growth.plugin.*

/**
 * Aplicação growth, com Kotlin :D
 *
 * @author Ariel Reis
 */

fun main() {
    embeddedServer(Netty, port = 8080, host = "0.0.0.0") {
        configureRouting()
        configureSerialization()
    }.start(wait = true)
}
