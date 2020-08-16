package br.dev.kurtis.domain.service;

import br.dev.kurtis.domain.model.Championship;
import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Matches;

public interface MatchService {
    Matches findAll();

    Match findOne(Long id);

    Championship findChampionship(Long id);
}
