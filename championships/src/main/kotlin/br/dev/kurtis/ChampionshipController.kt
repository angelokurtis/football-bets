package br.dev.kurtis

import br.dev.kurtis.model.Championship
import br.dev.kurtis.model.Resources
import com.fasterxml.jackson.databind.ObjectMapper
import io.micronaut.http.HttpHeaders
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import io.micronaut.http.hateoas.JsonError
import io.micronaut.http.hateoas.Link
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import java.io.File

@Controller("/championships")
class ChampionshipController(mapper: ObjectMapper, config: ChampionshipConfiguration) {

    private val resources: Resources
    private val log: Logger = LoggerFactory.getLogger(ChampionshipController::class.java)

    init {
        val content = File("${config.directory}/championships.json").readText()
        this.resources = mapper.readValue(content, Resources::class.java)
    }

    @Get(uri = "/", produces = ["application/hal+json"])
    fun findAll(headers: HttpHeaders): Resources {
        log.debug("${Trace(headers)}received request to get all championships")
        return resources
    }

    @Get(uri = "/{id}", produces = ["application/hal+json"])
    fun findOne(headers: HttpHeaders, id: Long): Championship? {
        log.debug("${Trace(headers)}received request to get the championship $id")
        val found = resources.embedded?.championships?.first { championship -> championship.hasId(id) }
        if (found == null) {
            log.debug("${Trace(headers)}the championship $id was not found")
        } else {
            log.debug("${Trace(headers)}the championship ${found.name} was found")
        }
        return found
    }

    @Error(exception = java.util.NoSuchElementException::class)
    fun notFound(request: HttpRequest<*>): HttpResponse<JsonError> {
        val error = JsonError("Page Not Found").link(Link.SELF, Link.of(request.uri))
        return HttpResponse.notFound(error)
    }
}


class Trace(headers: HttpHeaders) {
    private val traceId: String? = headers.get("X-B3-TraceId")
    private val spanId: String? = headers.get("X-B3-SpanId")
    private val sampled: String? = headers.get("X-B3-Sampled")

    override fun toString(): String {
        return if (isBlankString(traceId) && isBlankString(spanId) && isBlankString(sampled)) {
            ""
        } else String.format("[%s,%s,%s] ", traceId, spanId, sampled)
    }

    private fun isBlankString(string: String?): Boolean {
        return string == null || string.trim { it <= ' ' }.isEmpty()
    }
}