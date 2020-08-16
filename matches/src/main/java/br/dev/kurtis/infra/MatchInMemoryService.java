package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Link;
import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.MatchRelationships;
import br.dev.kurtis.domain.model.Matches;
import br.dev.kurtis.domain.service.MatchService;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.List;

@ApplicationScoped
public class MatchInMemoryService implements MatchService {

    private final Matches matches;
    private final List<MatchRelationships> relationships;

    @Inject
    public MatchInMemoryService(Matches matches, List<MatchRelationships> relationships) {
        this.matches = matches;
        this.relationships = relationships;
    }

    @Override
    public Matches findAll() {
        return this.matches;
    }

    @Override
    public Match findOne(Long id) {
        return this.matches.stream()
                .filter(match -> match.hasId(id))
                .findAny()
                .orElse(null);
    }

    @Override
    public Match findChampionship(Long id) {
        final var relationships = this.relationships.stream()
                .filter(relationship -> relationship.ofMatch(id))
                .findAny()
                .map(MatchRelationships::getChampionship)
                .map(Link::getHref);
        return null;
    }
}
