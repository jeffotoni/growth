package com.growth

import io.ktor.http.*
import kotlin.test.*
import io.ktor.server.testing.*
import com.growth.plugin.*

class ApplicationTest {
    @Test
    fun testRoot() {
        withTestApplication({ configureRouting() }) {
            handleRequest(HttpMethod.Get, "/").apply {
                assertEquals(HttpStatusCode.OK, response.status())
                assertEquals("Kotlin rocks \uD83D\uDE0A!", response.content)
            }
        }
    }
}