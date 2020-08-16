package br.dev.kurtis.domain.service;

import br.dev.kurtis.domain.model.Championship;
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;
import org.jboss.resteasy.annotations.jaxrs.PathParam;

import javax.ws.rs.GET;
import javax.ws.rs.Path;
import javax.ws.rs.Produces;

@Path("/")
@RegisterRestClient
public interface ChampionshipClient {
    @GET
    @Path("/{href}")
    @Produces("application/hal+json")
    Championship find(@PathParam String href);
}
