package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.*;
import br.dev.kurtis.domain.service.ChampionshipClient;
import br.dev.kurtis.domain.service.MatchService;
import org.eclipse.microprofile.rest.client.inject.RestClient;

import javax.enterprise.context.ApplicationScoped;
import javax.inject.Inject;
import java.util.List;

@ApplicationScoped
public class MatchInMemoryService implements MatchService {

    private final Matches matches;
    private final List<MatchRelationships> relationships;
    private final ChampionshipClient championshipClient;

    @Inject
    public MatchInMemoryService(Matches matches,
                                List<MatchRelationships> relationships,
                                @RestClient ChampionshipClient championshipClient) {
        this.matches = matches;
        this.relationships = relationships;
        this.championshipClient = championshipClient;
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
    public Championship findChampionship(Long id) {
        return this.relationships.stream()
                .filter(relationship -> relationship.ofMatch(id))
                .findAny()
                .map(MatchRelationships::getChampionship)
                .map(Link::getHref)
                .map(championshipClient::find)
                .map(championship -> {
                    championship.setSelf("/matches/" + id + "/championships");
                    return championship;
                })
                .orElse(null);
    }
}
