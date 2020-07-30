package br.dev.kurtis.infra;

import javax.enterprise.context.Dependent;
import javax.enterprise.inject.Produces;
import javax.json.bind.Jsonb;
import javax.json.bind.JsonbBuilder;

@Dependent
public class JsonConfiguration {
    @Produces
    public Jsonb jsonBinding() {
        return JsonbBuilder.create();
    }
}
