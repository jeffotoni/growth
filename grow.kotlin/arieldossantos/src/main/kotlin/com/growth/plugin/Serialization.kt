package com.growth.plugin

import io.ktor.jackson.*
import com.fasterxml.jackson.databind.*
import io.ktor.features.*
import io.ktor.application.*

fun Application.configureSerialization() {
    install(ContentNegotiation) {
        jackson {
                enable(SerializationFeature.INDENT_OUTPUT)
            }
    }
}
