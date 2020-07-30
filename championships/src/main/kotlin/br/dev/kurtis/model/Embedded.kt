package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.JsonProperty

data class Embedded(
        @JsonProperty("championships") var championships: List<Championship>? = null
)