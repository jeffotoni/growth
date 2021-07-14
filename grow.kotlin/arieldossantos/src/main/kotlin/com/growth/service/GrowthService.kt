package com.growth.service

import com.growth.response.GrowthListCountResponse
import com.growth.response.GrowthPostResponse
import com.growth.response.GrowthSizeResponse
import com.growth.response.GrowthValueResponse
import com.growth.storage.GrowthEntity
import io.ktor.util.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.GlobalScope
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import org.slf4j.LoggerFactory
import java.util.concurrent.atomic.AtomicBoolean


class GrowthService {
    companion object {
        private val logger = LoggerFactory.getLogger(GrowthService::class.java)
    }
    private var isProcessing = AtomicBoolean(false)

    fun insertValues(growthHashMapList: List<LinkedHashMap<String, Any>>): GrowthPostResponse {
        isProcessing.set(true) // define a variável is processing como true para verificaçao de status
        GlobalScope.launch { // cria um escopo de coroutine, para execução async =]
            growthHashMapList.forEach {
                safeInsert(it) //insere dados de forma segura (para não quebrar o processamento)
            }
            isProcessing.set(false) // define a variável is processing como false para verificaçao de status
        }
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

    fun findByCountryAndIndicatorAndYear(country: String, indicator: String, year: Int): GrowthEntity? {
        return GrowthEntity.growthStorage.firstOrNull {
            (it.Country.uppercase() == country.uppercase())
                .and(it.Indicator.uppercase() == indicator.uppercase())
                .and(it.Year == year)
        }
    }

    fun deleteByCountryAndIndicatorAndYear(country: String, indicator: String, year: Int): Boolean {
        val element = findByCountryAndIndicatorAndYear(country, indicator, year)
        if (element != null) {
            return GrowthEntity.growthStorage.remove(element)
        }
        return false
    }

    fun updateByCountryAndIndicatorAndYear(
        country: String, indicator: String, year: Int, value: Double
    ) {
        val currentElement = findByCountryAndIndicatorAndYear(country, indicator, year)
        if (currentElement != null) {
            val elementIndex = GrowthEntity.growthStorage.indexOf(currentElement)
            GrowthEntity.growthStorage[elementIndex] = currentElement.copy(Value = value)
        } else {
            GrowthEntity.growthStorage.add(
                GrowthEntity(
                    country,
                    indicator,
                    value,
                    year
                )
            )
        }
    }

    private fun getCountMessageAsString(value: Boolean): String {
        return if (!value) "complete" else "in progress"
    }

    private fun getTestValue(): Double? {
        return GrowthEntity.growthStorage.elementAtOrNull(0)?.Value
    }

    private fun getCountValue(): Int {
        return GrowthEntity.growthStorage.size
    }

    private fun safeInsert(growthHashMap: LinkedHashMap<String, Any>) {
        try {
            val currentGrowthEntity = GrowthEntity(
                growthHashMap["Country"].toString(),
                growthHashMap["Indicator"].toString(),
                growthHashMap["Value"].toString().toDouble(),
                growthHashMap["Year"].toString().toInt()
            )
            GrowthEntity.growthStorage.add(currentGrowthEntity)
        } catch (ex: Exception) {
            logger.error("Error on insert current item: $growthHashMap", ex)
        }
    }
}