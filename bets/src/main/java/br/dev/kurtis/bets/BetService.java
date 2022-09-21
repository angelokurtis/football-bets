package br.dev.kurtis.bets;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.http.HttpEntity;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import org.springframework.web.server.ResponseStatusException;

import java.security.SecureRandom;
import java.util.Optional;

@Service
public class BetService {
    private final RandomMatch randomMatch;
    private final RestTemplate matchesRestTemplate;
    private final RestTemplate teamsRestTemplate;

    public BetService(
            final RandomMatch randomMatch,
            @Qualifier("matchesRestTemplate") final RestTemplate matchesRestTemplate,
            @Qualifier("teamsRestTemplate") final RestTemplate teamsRestTemplate) {
        this.randomMatch = randomMatch;
        this.matchesRestTemplate = matchesRestTemplate;
        this.teamsRestTemplate = teamsRestTemplate;
    }

    public Bet betRandomly() {
        final Optional<Match> match = this.randomMatch.findOne();
        final Team homeTeam = match.map(Match::getScoreHome)
                .flatMap(this::findTeamFromScore)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.SERVICE_UNAVAILABLE, "teams service is unavailable"));
        final Team awayTeam = match.map(Match::getScoreAway)
                .flatMap(this::findTeamFromScore)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.SERVICE_UNAVAILABLE, "teams service is unavailable"));
        final Championship championship = match.map(Match::getLinks)
                .map(MatchLinks::getChampionship)
                .map(Link::getHref)
                .map(href -> this.matchesRestTemplate.getForEntity(href, Championship.class))
                .filter(HttpEntity::hasBody)
                .map(HttpEntity::getBody)
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.SERVICE_UNAVAILABLE, "championships service is unavailable"));
        return Bet.instanceOf(match, homeTeam, awayTeam, championship);
    }

    private Optional<Team> findTeamFromScore(final Score score) {
        return Optional.of(score)
                .map(Score::getLinks)
                .map(ScoreLinks::getTeam)
                .map(Link::getHref)
                .map(href -> this.teamsRestTemplate.getForEntity(href, Team.class))
                .filter(HttpEntity::hasBody)
                .map(HttpEntity::getBody);
    }


}
