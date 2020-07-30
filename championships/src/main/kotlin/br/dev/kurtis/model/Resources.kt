package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.JsonProperty

data class Resources(
        @JsonProperty("_embedded") var embedded: Embedded? = null,
        @JsonProperty("_links") var links: Links? = null
)