package br.dev.kurtis.domain.model;

public class MatchLinks {

    private Link self;
    private Link match;
    private Link championship;

    public Link getSelf() {
        return self;
    }

    public void setSelf(Link self) {
        this.self = self;
    }

    public Link getMatch() {
        return match;
    }

    public void setMatch(Link match) {
        this.match = match;
    }

    public Link getChampionship() {
        return championship;
    }

    public void setChampionship(Link championship) {
        this.championship = championship;
    }
}
