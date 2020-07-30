package br.dev.kurtis.model

import com.fasterxml.jackson.annotation.JsonProperty

data class Links(
        @JsonProperty("championship") var championship: Link? = null,
        @JsonProperty("self") var self: Link? = null,
        @JsonProperty("teams") var teams: Link? = null
)