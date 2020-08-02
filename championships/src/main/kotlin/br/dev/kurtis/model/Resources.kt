package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.*
import io.micronaut.core.annotation.Introspected
import java.util.*

@Introspected
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder("_embedded", "_links")
class Resources {
    @get:JsonProperty("_embedded")
    @set:JsonProperty("_embedded")
    @JsonProperty("_embedded")
    var embedded: Embedded? = null

    @get:JsonProperty("_links")
    @set:JsonProperty("_links")
    @JsonProperty("_links")
    var links: Link? = null

    @JsonIgnore
    private val additionalProperties: MutableMap<String, Any> = HashMap()

    @JsonAnyGetter
    fun getAdditionalProperties(): Map<String, Any> = additionalProperties

    @JsonAnySetter
    fun setAdditionalProperty(name: String, value: Any) {
        additionalProperties[name] = value
    }
}