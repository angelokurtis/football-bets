
package br.dev.kurtis.domain.model;

import javax.json.bind.annotation.JsonbProperty;
import java.util.Optional;

public class Championship {

    @JsonbProperty("_links")
    private ChampionshipLinks links;
    private String name;
    private Integer year;

    public ChampionshipLinks getLinks() {
        return links;
    }

    public void setLinks(ChampionshipLinks links) {
        this.links = links;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public Integer getYear() {
        return year;
    }

    public void setYear(Integer year) {
        this.year = year;
    }

    public void setSelf(String self) {
        Optional.of(this.links)
                .map(ChampionshipLinks::getSelf)
                .ifPresent(link -> link.setHref(self));
    }
}
