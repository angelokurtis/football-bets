package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.JsonProperty

data class Championship(
        @JsonProperty("_links") var links: Links? = null,
        @JsonProperty("name") var name: String? = null,
        @JsonProperty("year") var year: Int? = null
)
{
    fun hasId(id: Long): Boolean {
        return this.links?.self?.href.let { self -> self == "/championships/$id" }
    }
}
