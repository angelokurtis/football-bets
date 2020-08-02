package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.*
import io.micronaut.core.annotation.Introspected
import java.util.*

@Introspected
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder("href")
class Link {
    @get:JsonProperty("href")
    @set:JsonProperty("href")
    @JsonProperty("href")
    var href: String? = null

    @JsonIgnore
    private val additionalProperties: MutableMap<String, Any> = HashMap()

    @JsonAnyGetter
    fun getAdditionalProperties(): Map<String, Any> = additionalProperties

    @JsonAnySetter
    fun setAdditionalProperty(name: String, value: Any) {
        additionalProperties[name] = value
    }
}