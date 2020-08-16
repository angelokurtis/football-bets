package br.dev.kurtis.domain.model;

import javax.json.bind.annotation.JsonbProperty;
import java.util.Collection;
import java.util.Optional;
import java.util.stream.Stream;

public class Matches {
    @JsonbProperty("_embedded")
    private Embedded embedded;
    @JsonbProperty("_links")
    private MatchesLinks links;

    public Stream<Match> stream() {
        return Optional.ofNullable(this.embedded)
                .map(Embedded::getMatches)
                .stream()
                .flatMap(Collection::stream);
    }

    public Embedded getEmbedded() {
        return embedded;
    }

    public void setEmbedded(Embedded embedded) {
        this.embedded = embedded;
    }

    public MatchesLinks getLinks() {
        return links;
    }

    public void setLinks(MatchesLinks links) {
        this.links = links;
    }
}
