package br.dev.kurtis.domain.model;

import javax.json.bind.annotation.JsonbProperty;

public class ScoreHome {

    private Integer goals;
    @JsonbProperty("_links")
    private ScoreLinks links;

    public Integer getGoals() {
        return goals;
    }

    public void setGoals(Integer goals) {
        this.goals = goals;
    }

    public ScoreLinks getLinks() {
        return links;
    }

    public void setLinks(ScoreLinks links) {
        this.links = links;
    }
}
