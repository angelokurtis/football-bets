package br.dev.kurtis.infra;

import br.dev.kurtis.domain.model.Championship;
import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Matches;
import br.dev.kurtis.domain.service.MatchService;
import org.jboss.logging.Logger;

import javax.inject.Inject;
import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.PathParam;
import javax.ws.rs.Produces;
import javax.ws.rs.core.Context;
import javax.ws.rs.core.HttpHeaders;
import javax.ws.rs.core.MediaType;

@Path("/matches")
public class MatchController {

    private static final Logger LOG = Logger.getLogger(MatchController.class);
    private final MatchService service;

    @Inject
    public MatchController(MatchService service) {
        this.service = service;
    }

    @GET
    @Produces(MediaType.APPLICATION_JSON)
    public Matches findAll(@Context HttpHeaders headers) {
        LOG.info(new Trace(headers).toString() + "received request for all matches");
        return this.service.findAll();
    }

    @GET
    @Path("/{id}")
    @Produces(MediaType.APPLICATION_JSON)
    public Match findOne(@Context HttpHeaders headers, @PathParam(value = "id") Long id) {
        LOG.infof(new Trace(headers).toString() + "received request for match '%s'", id);
        return this.service.findOne(id);
    }

    @GET
    @Path("/{id}/championship")
    @Produces(MediaType.APPLICATION_JSON)
    public Championship findChampionship(@Context HttpHeaders headers, @PathParam(value = "id") Long id) {
        LOG.infof(new Trace(headers).toString() + "received request for the championship of match '%s'", id);
        return this.service.findChampionship(id);
    }
}