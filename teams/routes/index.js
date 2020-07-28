const express = require('express');
const fs = require('fs');

const teams = JSON.parse(fs.readFileSync(`${process.env.DATA_DIR}/teams.json`));
const router = express.Router();

const getTeams = async (req, res, next) => {
    try {
        res.json(teams);
    } catch (e) {
        next(e);
    }
};

const getTeam = async (req, res, next) => {
    try {
        const team = teams._embedded.teams.filter(({_links}) => _links.team.href === `/teams/${req.params.id}`)
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

module.exports = router;