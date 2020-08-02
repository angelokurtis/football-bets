package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.*
import io.micronaut.core.annotation.Introspected
import java.util.*

@Introspected
@JsonInclude(JsonInclude.Include.NON_NULL)
@JsonPropertyOrder("self", "championship", "teams")
class Links {
    @get:JsonProperty("self")
    @set:JsonProperty("self")
    @JsonProperty("self")
    var self: Link? = null

    @get:JsonProperty("championship")
    @set:JsonProperty("championship")
    @JsonProperty("championship")
    var championship: Link? = null

    @get:JsonProperty("teams")
    @set:JsonProperty("teams")
    @JsonProperty("teams")
    var teams: Link? = null

    @JsonIgnore
    private val additionalProperties: MutableMap<String, Any> = HashMap()

    @JsonAnyGetter
    fun getAdditionalProperties(): Map<String, Any> = additionalProperties

    @JsonAnySetter
    fun setAdditionalProperty(name: String, value: Any) {
        additionalProperties[name] = value
    }
}