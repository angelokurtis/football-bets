
package br.dev.kurtis.domain.model;

public class ChampionshipLinks {

    private Link self;
    private Link championship;
    private Link teams;

    public Link getSelf() {
        return self;
    }

    public void setSelf(Link self) {
        this.self = self;
    }

    public Link getChampionship() {
        return championship;
    }

    public void setChampionship(Link championship) {
        this.championship = championship;
    }

    public Link getTeams() {
        return teams;
    }

    public void setTeams(Link teams) {
        this.teams = teams;
    }
}
