package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Resource;

import javax.enterprise.context.Dependent;
import javax.enterprise.inject.Produces;
import javax.json.bind.Jsonb;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

@Dependent
public class MatchConfiguration {

    @Produces
    public Resource matchesFromJson(Jsonb jsonb) throws IOException {
        Path path = Path.of("/home/kurtis/dev/projects/labs/football-bets-api/responses/matches.json");
        String content = Files.readString(path);

        return jsonb.fromJson(content, Resource.class);
    }
}
