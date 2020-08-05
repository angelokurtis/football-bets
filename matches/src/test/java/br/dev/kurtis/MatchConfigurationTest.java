package br.dev.kurtis;

import br.dev.kurtis.infra.MatchConfiguration;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import javax.json.bind.JsonbBuilder;
import java.io.IOException;

class MatchConfigurationTest {

    @Test
    @DisplayName("Load matches from JSON")
    void loadMatchesFromJson() throws IOException {
        final var jsonb = JsonbBuilder.create();
        new MatchConfiguration("/home/kurtis/dev/projects/labs/football-bets-api/responses").matchesFromJson(jsonb);
    }
}