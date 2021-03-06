const express = require('express');
const fs = require('fs');

const teams = JSON.parse(fs.readFileSync(`${process.env.DATA_DIR}/teams.json`));
const router = express.Router();

class TeamV1 {
    constructor({name, _links}) {
        this.name = name;
        this._links = _links
    }
}

class TeamV2 {
    constructor({name, image_url, _links}) {
        this.name = name;
        this.image_url = image_url
        this._links = _links
    }
}

const Team = process.env.VERSION === 'v2' ? TeamV2 : TeamV1

const getTeams = async (req, res, next) => {
    console.log(new Trace(req).string() + "received request to get all teams")
    try {
        res.json(teams);
    } catch (e) {
        next(e);
    }
};

const getTeam = async (req, res, next) => {
    console.log(new Trace(req).string() + "received request to get the team " + req.params.id)
    try {
        const team = teams._embedded.teams
            .filter(({_links}) => _links.team.href === `/teams/${req.params.id}`)
            .map(value => new Team(value))
        if (!team || team.length === 0) {
            const err = new Error('Team not found');
            err.status = 404;
            throw err;
        }
        res.json(team[0]);
    } catch (e) {
        next(e);
    }
};

router
    .route('/teams')
    .get(getTeams);
router
    .route('/teams/:id')
    .get(getTeam);

class Trace {
    constructor(req) {
        this.traceId = req.header('X-B3-TraceId')
        this.spanId = req.header('X-B3-SpanId')
        this.sampled = req.header('X-B3-Sampled')
    }

    string() {
        if (this.traceId || this.spanId || this.sampled) {
            return `[${this.traceId},${this.spanId},${this.sampled}] `
        }
        return ""
    }
}

module.exports = router;