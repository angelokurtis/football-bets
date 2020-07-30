package br.dev.kurtis.domain.model;

import java.util.List;

public class Embedded {

    private List<Match> matches;

    public List<Match> getMatches() {
        return matches;
    }

    public void setMatches(List<Match> matches) {
        this.matches = matches;
    }
}
