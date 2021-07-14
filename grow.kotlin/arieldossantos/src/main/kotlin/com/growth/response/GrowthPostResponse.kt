package com.growth.response

import com.fasterxml.jackson.annotation.JsonInclude

@JsonInclude(JsonInclude.Include.NON_NULL)
data class GrowthPostResponse(
    val msg: String
)