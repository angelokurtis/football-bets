package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Resource;
import br.dev.kurtis.domain.service.MatchService;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;

@ApplicationScoped
public class MatchInMemoryService implements MatchService {

    private final Resource resource;

    @Inject
    public MatchInMemoryService(Resource resource) {
        this.resource = resource;
    }

    @Override
    public Resource findAll() {
        return this.resource;
    }

    @Override
    public Match findOne(Long id) {
        return this.resource.stream()
                .filter(match -> match.hasId(id))
                .findAny()
                .orElse(null);
    }
}
