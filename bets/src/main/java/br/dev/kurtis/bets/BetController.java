package br.dev.kurtis.bets;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class BetController {
    private final BetService service;

    public BetController(final BetService service) {
        this.service = service;
    }

    @PostMapping(path = "/bets", produces = "application/hal+json")
    private Bet create() {
        return this.service.betRandomly();
    }
}
