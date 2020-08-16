package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Matches;
import br.dev.kurtis.domain.service.MatchService;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;

@ApplicationScoped
public class MatchInMemoryService implements MatchService {

    private final Matches matches;

    @Inject
    public MatchInMemoryService(Matches matches) {
        this.matches = matches;
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
}
