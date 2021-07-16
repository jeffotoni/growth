package com.growth.service

import com.growth.response.GrowthListCountResponse
import com.growth.response.GrowthPostResponse
import com.growth.response.GrowthSizeResponse
import com.growth.storage.GrowthEntity
import org.slf4j.LoggerFactory
import java.util.*
import java.util.concurrent.atomic.AtomicBoolean


class GrowthService {
    companion object {
        private val logger = LoggerFactory.getLogger(GrowthService::class.java)
    }
    private var isProcessing = AtomicBoolean(false)

    fun insertValues(growthHashMapList: List<LinkedHashMap<String, Any>>): GrowthPostResponse {
        isProcessing.set(true) // define a variável is processing como true para verificaçao de status
        growthHashMapList.forEach {
            safeInsert(it) //insere dados de forma segura (para não quebrar o processamento)
        }
        isProcessing.set(false) // define a variável is processing como false para verificaçao de status
        return GrowthPostResponse("In progress")
    }

    fun getSize(): GrowthSizeResponse {
        return GrowthSizeResponse(
            getCountValue()
        )
    }

    fun getGrowthEntityListSize(): GrowthListCountResponse {
        val currentProcessing = isProcessing.get()
        return GrowthListCountResponse(
            getCountMessageAsString(currentProcessing),
            getTestValue(),
            getCountValue()
        )
    }

    @Synchronized fun findByCountryAndIndicatorAndYear(country: String, indicator: String, year: Int): GrowthEntity? {
        return GrowthEntity.growthStorage.filter {
            (it.value.Country.uppercase() == country.uppercase())
                .and(it.value.Indicator.uppercase() == indicator.uppercase())
                .and(it.value.Year == year)
        }.values.firstOrNull()
    }

    @Synchronized fun deleteByCountryAndIndicatorAndYear(country: String, indicator: String, year: Int): Boolean {
        val key = getKeyByCountryAndIndicatorAndYear(country, indicator, year) ?: UUID.randomUUID()
        if (key != null) {
            val element = GrowthEntity.growthStorage[key]
            return GrowthEntity.growthStorage.remove(key, element)
        }
        return false
    }

    fun updateByCountryAndIndicatorAndYear(
        country: String, indicator: String, year: Int, value: Double
    ) {
        val currentKey = getKeyByCountryAndIndicatorAndYear(country, indicator, year) ?: UUID.randomUUID()

        GrowthEntity.growthStorage[currentKey] = GrowthEntity(
            country,
            indicator,
            value,
            year
        )
    }

    private fun getKeyByCountryAndIndicatorAndYear(
        country: String, indicator: String, year: Int
    ): UUID? {
        return GrowthEntity.growthStorage.filterValues {
            it.Country == country && it.Indicator == indicator && it.Year == year
        }.keys.firstOrNull()
    }

    private fun getCountMessageAsString(value: Boolean): String {
        return if (!value) "complete" else "in progress"
    }

    private fun getTestValue(): Double? {
        val treeMap = TreeMap(GrowthEntity.growthStorage)
        return if(treeMap.isNotEmpty()) treeMap.firstEntry().value.Value else null
    }

    private fun getCountValue(): Int {
        return if(GrowthEntity.growthStorage.isNotEmpty())
            GrowthEntity.growthStorage.size
        else
            0
    }

    private fun safeInsert(growthHashMap: LinkedHashMap<String, Any>) {
        try {
            val currentGrowthEntity = GrowthEntity(
                growthHashMap["Country"].toString(),
                growthHashMap["Indicator"].toString(),
                growthHashMap["Value"].toString().toDouble(),
                growthHashMap["Year"].toString().toInt()
            )
            GrowthEntity.growthStorage[UUID.randomUUID()] = currentGrowthEntity
        } catch (ex: Exception) {
            logger.error("Error on insert current item: $growthHashMap", ex)
        }
    }
}