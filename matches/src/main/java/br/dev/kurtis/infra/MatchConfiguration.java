package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.MatchRelationships;
import br.dev.kurtis.domain.model.Matches;
import org.eclipse.microprofile.config.inject.ConfigProperty;

import javax.enterprise.context.Dependent;
import javax.enterprise.inject.Produces;
import javax.json.bind.Jsonb;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.List;

@Dependent
public class MatchConfiguration {

    private final String directory;

    public MatchConfiguration(@ConfigProperty(name = "matches.directory") String directory) {
        this.directory = directory;
    }

    @Produces
    public Matches matchesFromJson(Jsonb jsonb) throws IOException {
        Path path = Path.of(this.directory + "/matches.json");
        String content = Files.readString(path);

        return jsonb.fromJson(content, Matches.class);
    }

    @Produces
    public List<MatchRelationships> relationshipsFromJson(Jsonb jsonb) throws IOException {
        Path path = Path.of(this.directory + "/relationships.json");
        String content = Files.readString(path);

        return jsonb.fromJson(content, new ArrayList<MatchRelationships>(){}.getClass().getGenericSuperclass());
    }
}
