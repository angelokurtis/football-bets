package br.dev.kurtis.domain.service;

import br.dev.kurtis.domain.model.Match;
import br.dev.kurtis.domain.model.Resource;

public interface MatchService {
    Resource findAll();

    Match findOne(Long id);
}
