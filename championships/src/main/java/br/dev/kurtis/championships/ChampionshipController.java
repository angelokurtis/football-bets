package br.dev.kurtis.championships;

import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.util.Collection;
import java.util.Optional;

@RestController
@AllArgsConstructor
public class ChampionshipController {
    private final ChampionshipService service;

    @GetMapping(path = "/championships", produces = "application/hal+json")
    private Championships findAll() {
        return this.service.deserializeChampionshipsJSON();
    }

    @GetMapping(path = "/championships/{id}", produces = "application/hal+json")
    private Optional<Championship> findOne(@PathVariable("id") final String id) {
        final Championships championships = this.service.deserializeChampionshipsJSON();
        return Optional.ofNullable(championships)
                .map(Championships::getEmbedded)
                .map(Embedded::getChampionships).stream()
                .flatMap(Collection::stream)
                .filter(team -> Optional.of(team)
                        .map(Championship::getLinks)
                        .map(Links::getChampionship)
                        .map(Link::getHref)
                        .orElse("")
                        .equals("/championships/" + id))
                .findFirst();
    }
}
