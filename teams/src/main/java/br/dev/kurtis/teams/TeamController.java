package br.dev.kurtis.teams;

import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.Optional;

@RestController
@AllArgsConstructor
public class TeamController {
    private final TeamService service;

    @GetMapping(path = "/teams", produces = "application/hal+json")
    private Teams findAll() {
        return this.service.deserializeTeamsJSON();
    }

    @GetMapping(path = "/teams/{id}", produces = "application/hal+json")
    private Optional<Team> findOne(@PathVariable("id") final String id) {
        final Teams teams = this.service.deserializeTeamsJSON();
        return Optional.ofNullable(teams)
                .map(Teams::getEmbedded)
                .map(Embedded::getTeams).stream()
                .flatMap(Collection::stream)
                .filter(team -> Optional.of(team)
                        .map(Team::getLinks)
                        .map(Links::getTeam)
                        .map(Link::getHref)
                        .orElse("")
                        .equals("/teams/" + id))
                .findFirst();
    }
}
