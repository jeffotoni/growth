package com.growth.storage

import com.fasterxml.jackson.annotation.JsonProperty
import com.fasterxml.jackson.databind.PropertyNamingStrategy
import com.fasterxml.jackson.databind.annotation.JsonNaming

@JsonNaming(PropertyNamingStrategy.UpperCamelCaseStrategy::class)
data class GrowthEntity(
    @JsonProperty("Country")
    var Country: String,
    @JsonProperty("Indicator")
    var Indicator: String,
    @JsonProperty("Value")
    var Value: Double,
    @JsonProperty("Year")
    var Year: Int
) {
    companion object {
        val growthStorage = mutableListOf<GrowthEntity>()
    }
}
