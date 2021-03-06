package br.dev.kurtis;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

import static io.restassured.RestAssured.given;
import static org.hamcrest.CoreMatchers.is;

@QuarkusTest
public class MatchControllerTest {

    @Test
    public void testHelloEndpoint() {
        given()
          .when().get("/api/matches")
          .then()
             .statusCode(200);
    }

}