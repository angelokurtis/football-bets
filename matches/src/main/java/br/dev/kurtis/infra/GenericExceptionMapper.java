package br.dev.kurtis.infra;

import org.jboss.logging.Logger;

import javax.ws.rs.core.Context;
import javax.ws.rs.core.HttpHeaders;
import javax.ws.rs.core.Response;
import javax.ws.rs.ext.ExceptionMapper;
import javax.ws.rs.ext.Provider;

@Provider
public class GenericExceptionMapper implements ExceptionMapper<Throwable> {

    private static final Logger LOG = Logger.getLogger(GenericExceptionMapper.class);
    @Context
    private HttpHeaders headers;

    @Override
    public Response toResponse(Throwable throwable) {
        LOG.error(new Trace(headers).toString() + "unhandled error", throwable);
        ErrorResponse o = new ErrorResponse("The server has encountered a situation it doesn't know how to handle.", 500);
        return Response.serverError().entity(o).build();
    }
}

