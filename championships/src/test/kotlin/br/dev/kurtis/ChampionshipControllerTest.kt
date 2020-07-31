package br.dev.kurtis

import io.micronaut.http.HttpStatus
import io.micronaut.http.client.RxHttpClient
import io.micronaut.http.client.exceptions.HttpClientResponseException
import io.micronaut.runtime.server.EmbeddedServer
import io.micronaut.test.annotation.MicronautTest
import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

@MicronautTest
class ChampionshipControllerTest(private val server: EmbeddedServer) {

    @Test
    fun `Server is running`() {
        assert(server.isRunning)
    }

    @Test
    fun `Found championship by ID`() {
        val client: RxHttpClient = server.applicationContext.createBean(RxHttpClient::class.java, server.url)
        val response = client.toBlocking().exchange("/championships/924", String::class.java)
        assertEquals(HttpStatus.OK, response.status())
        println(response.body())
        client.close()
    }

    @Test
    fun `Not found championship by ID`() {
        val client: RxHttpClient = server.applicationContext.createBean(RxHttpClient::class.java, server.url)
        val response = try {
            client.toBlocking().exchange("/championships/925", String::class.java)
        } catch (ex: HttpClientResponseException) {
            ex.response
        } finally {
            client.close()
        }
        assertEquals(HttpStatus.NOT_FOUND, response.status())
        println(response.body())
    }
}