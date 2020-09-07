package br.dev.kurtis.domain.service;

import br.dev.kurtis.domain.model.Championship;
import org.eclipse.microprofile.rest.client.inject.RegisterRestClient;
import org.jboss.resteasy.annotations.jaxrs.HeaderParam;
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
    Championship find(@HeaderParam("X-B3-TraceId") String traceId,
                      @HeaderParam("X-B3-ParentSpanId") String parentSpanId,
                      @HeaderParam("X-B3-SpanId") String spanId,
                      @HeaderParam("X-B3-Sampled") String sampled,
                      @PathParam String href);
}
