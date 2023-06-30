package br.dev.kurtis.championships;

import com.fasterxml.jackson.annotation.JsonInclude;
import com.fasterxml.jackson.annotation.JsonProperty;

@lombok.Data
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Championship {
    @JsonProperty("name")
    private String name;
    @JsonProperty("year")
    private Long year;
    @JsonProperty("_links")
    private Links links;
}