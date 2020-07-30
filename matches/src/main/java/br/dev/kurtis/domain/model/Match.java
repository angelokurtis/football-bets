package br.dev.kurtis.domain.model;

import javax.json.bind.annotation.JsonbProperty;
import java.util.Optional;

public class Match {

    private String date;
    private String status;
    @JsonbProperty("score_home")
    private ScoreHome scoreHome;
    @JsonbProperty("score_away")
    private ScoreAway scoreAway;
    @JsonbProperty("_links")
    private MatchLinks links;

    public boolean hasId(Long id) {
        return Optional.ofNullable(this.getLinks())
                .map(MatchLinks::getSelf)
                .map(Link::getHref)
                .filter(self -> self.equals("/matches/" + id))
                .isPresent();
    }

    public String getDate() {
        return date;
    }

    public void setDate(String date) {
        this.date = date;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public ScoreHome getScoreHome() {
        return scoreHome;
    }

    public void setScoreHome(ScoreHome scoreHome) {
        this.scoreHome = scoreHome;
    }

    public ScoreAway getScoreAway() {
        return scoreAway;
    }

    public void setScoreAway(ScoreAway scoreAway) {
        this.scoreAway = scoreAway;
    }

    public MatchLinks getLinks() {
        return links;
    }

    public void setLinks(MatchLinks links) {
        this.links = links;
    }

}
