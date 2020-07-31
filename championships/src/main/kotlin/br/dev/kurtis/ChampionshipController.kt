package br.dev.kurtis

import br.dev.kurtis.model.Resources
import com.fasterxml.jackson.databind.ObjectMapper
import io.micronaut.http.HttpRequest
import io.micronaut.http.HttpResponse
import io.micronaut.http.annotation.Controller
import io.micronaut.http.annotation.Error
import io.micronaut.http.annotation.Get
import io.micronaut.http.hateoas.JsonError
import io.micronaut.http.hateoas.Link
import java.io.File

@Controller("/championships")
class ChampionshipController(mapper: ObjectMapper) {

    private val resources: Resources

    init {
        val content = File("/home/kurtis/dev/projects/labs/football-bets-api/responses/championships.json").readText()
        this.resources = mapper.readValue(content, Resources::class.java)
    }

    @Get(uri = "/", produces = ["application/hal+json"])
    fun findAll() = this.resources

    @Get(uri = "/{id}", produces = ["application/hal+json"])
    fun findOne(id: Long) = resources.embedded?.championships?.first { championship -> championship.hasId(id) }

    @Error(exception = java.util.NoSuchElementException::class)
    fun notFound(request: HttpRequest<*>): HttpResponse<JsonError> {
        val error = JsonError("Page Not Found").link(Link.SELF, Link.of(request.uri))
        return HttpResponse.notFound(error)
    }
}