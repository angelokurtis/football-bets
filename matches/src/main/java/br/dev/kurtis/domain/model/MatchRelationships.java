
package br.dev.kurtis.domain.model;

import java.util.Optional;

public class MatchRelationships {

    private Link self;
    private Link championship;

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

    public boolean ofMatch(Long id) {
        return Optional.ofNullable(this.getSelf())
                .map(Link::getHref)
                .filter(self -> self.equals("/matches/" + id))
                .isPresent();
    }
}
