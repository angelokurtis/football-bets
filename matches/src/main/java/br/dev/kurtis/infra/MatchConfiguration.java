package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Resource;
import org.eclipse.microprofile.config.inject.ConfigProperty;

import javax.enterprise.context.Dependent;
import javax.enterprise.inject.Produces;
import javax.json.bind.Jsonb;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;

@Dependent
public class MatchConfiguration {

    private final String directory;

    public MatchConfiguration(@ConfigProperty(name = "matches.directory") String directory) {
        this.directory = directory;
    }

    @Produces
    public Resource matchesFromJson(Jsonb jsonb) throws IOException {
        Path path = Path.of(this.directory + "/matches.json");
        String content = Files.readString(path);

        return jsonb.fromJson(content, Resource.class);
    }
}
