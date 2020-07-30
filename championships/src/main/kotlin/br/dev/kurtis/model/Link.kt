package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.JsonProperty

data class Link(
        @JsonProperty("href") var href: String? = null
)