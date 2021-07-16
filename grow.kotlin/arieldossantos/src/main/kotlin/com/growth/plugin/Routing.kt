package com.growth.plugin

import com.growth.request.GrowthValueRequest
import com.growth.service.GrowthService
import com.growth.storage.GrowthEntity
import io.ktor.routing.*
import io.ktor.locations.*
import io.ktor.application.*
import io.ktor.http.*
import io.ktor.request.*
import io.ktor.response.*

fun Application.configureRouting() {

    val growthService = GrowthService()

    install(Locations) {
    }

    routing {
        get("/") {
            call.respondText("Kotlin rocks \uD83D\uDE0A!")
        }

        get("/ping") {
            call.respond(HttpStatusCode.OK, "pong")
        }

        route("/api/v1/growth") {
            //Post on growth
            post() {
                val growthEntityList = call.receive<List<LinkedHashMap<String, Any>>>()
                call.respond(HttpStatusCode.Accepted, growthService.insertValues(growthEntityList))
            }

            get("/post/status") {
                call.respond(HttpStatusCode.OK, growthService.getGrowthEntityListSize())
            }

            get("/size") {
                call.respond(HttpStatusCode.OK, growthService.getSize())
            }

            route("/{country}/{indicator}/{year}") {
                get {
                    val country = call.parameters["country"]
                    val indicator = call.parameters["indicator"]
                    val year = call.parameters["year"]
                    country?.let {
                        indicator?.let {
                            year?.let {
                                call.respond(
                                    HttpStatusCode.OK,
                                    growthService.findByCountryAndIndicatorAndYear(
                                        country, indicator, year.toInt()
                                    ) ?: "{}"
                                )
                            }
                        }
                    }
                }

                put {
                    val country = call.parameters["country"]
                    val indicator = call.parameters["indicator"]
                    val year = call.parameters["year"]
                    val valueModel = call.receive<GrowthValueRequest>()
                    country?.let {
                        indicator?.let {
                            year?.let {
                                growthService.updateByCountryAndIndicatorAndYear(
                                    country, indicator, year.toInt(), valueModel.value
                                )
                                call.respond(
                                    HttpStatusCode.Accepted
                                )
                            }
                        }
                    }
                }

                delete {
                    val country = call.parameters["country"]
                    val indicator = call.parameters["indicator"]
                    val year = call.parameters["year"]
                    country?.let {
                        indicator?.let {
                            year?.let {
                                growthService.deleteByCountryAndIndicatorAndYear(
                                    country, indicator, year.toInt()
                                )
                                call.respond(
                                    HttpStatusCode.Accepted
                                )
                            }
                        }
                    }
                }

            }
        }
    }
}
