package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Resource;
import br.dev.kurtis.domain.service.MatchService;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.MediaType;

@Path("/api/matches")
public class MatchController {

    private final MatchService service;

    @Inject
    public MatchController(MatchService service) {
        this.service = service;
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public Resource findAll() {
        return this.service.findAll();
    }

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    public Match findOne(@PathParam(value = "id") Long id) {
        return this.service.findOne(id);
    }
}