const express = require('express');
const path = require('path');
const app = express();

const PORT = 3000;

app.set('port', PORT);
app.set('env', 'production');

app.use('/', require(path.join(__dirname, 'routes')));

app.use((req, res, next) => {
    const err = new Error(`${req.method} ${req.url} Not Found`);
    err.status = 404;
    next(err);
});

app.use((err, req, res, next) => {
    if (!err.status) console.error(err);
    res.status(err.status || 500);
    res.json({error: {message: err.message}});
});

app.listen(PORT, () => {
    console.log(`Express Server started on Port ${app.get('port')}`);
});