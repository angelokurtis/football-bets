package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.*
import io.micronaut.core.annotation.Introspected
import java.util.*

@Introspected
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder("name", "year", "_links")
class Championship {
    @get:JsonProperty("name")
    @set:JsonProperty("name")
    @JsonProperty("name")
    var name: String? = null

    @get:JsonProperty("year")
    @set:JsonProperty("year")
    @JsonProperty("year")
    var year: Int? = null

    @get:JsonProperty("_links")
    @set:JsonProperty("_links")
    @JsonProperty("_links")
    var links: Links? = null

    @JsonIgnore
    private val additionalProperties: MutableMap<String, Any> = HashMap()

    @JsonAnyGetter
    fun getAdditionalProperties(): Map<String, Any> = additionalProperties

    @JsonAnySetter
    fun setAdditionalProperty(name: String, value: Any) {
        additionalProperties[name] = value
    }

    fun hasId(id: Long) = this.links?.self?.href.let { self -> self == "/championships/$id" }
}