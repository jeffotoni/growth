package com.growth.response

data class GrowthListCountResponse(
    val msg: String,
    val testValue: Double? = null,
    val count: Int = 0
)
