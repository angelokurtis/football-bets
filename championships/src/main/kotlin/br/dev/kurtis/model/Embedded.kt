package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.*
import io.micronaut.core.annotation.Introspected
import java.util.*

@Introspected
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder("championships")
class Embedded {
    @get:JsonProperty("championships")
    @set:JsonProperty("championships")
    @JsonProperty("championships")
    var championships: List<Championship>? = null

    @JsonIgnore
    private val additionalProperties: MutableMap<String, Any> = HashMap()

    @JsonAnyGetter
    fun getAdditionalProperties(): Map<String, Any> = additionalProperties

    @JsonAnySetter
    fun setAdditionalProperty(name: String, value: Any) {
        additionalProperties[name] = value
    }
}