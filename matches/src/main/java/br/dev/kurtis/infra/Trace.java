package br.dev.kurtis.infra;

import javax.ws.rs.core.HttpHeaders;

public class Trace {
    private final String traceId;
    private final String parentSpanId;
    private final String spanId;
    private final String sampled;

    public Trace(HttpHeaders headers) {
        final var requestHeaders = headers.getRequestHeaders();
        this.traceId = requestHeaders.getFirst("X-B3-TraceId");
        this.parentSpanId = requestHeaders.getFirst("X-B3-ParentSpanId");
        this.spanId = requestHeaders.getFirst("X-B3-SpanId");
        this.sampled = requestHeaders.getFirst("X-B3-Sampled");
    }

    public String getTraceId() {
        return traceId;
    }

    public String getParentSpanId() {
        return parentSpanId;
    }

    public String getSpanId() {
        return spanId;
    }

    public String getSampled() {
        return sampled;
    }

    @Override
    public String toString() {
        if (isBlankString(this.traceId) && isBlankString(this.spanId) && isBlankString(this.sampled)) {
            return "";
        }
        return String.format("[%s,%s,%s] ", this.traceId, this.spanId, this.sampled);
    }

    private boolean isBlankString(String string) {
        return string == null || string.trim().isEmpty();
    }
}
