const http = require("http");
const os = require("os");

const PORT = process.env.PORT || 3000;
const MESSAGE = process.env.MESSAGE || "Harissh Ragav Dhamodaran 2023BCD0059";

const server = http.createServer((req, res) => {

    if (req.url === "/") {
        res.writeHead(200);
        res.end(JSON.stringify({
            message: MESSAGE,
            host: os.hostname()
        }));
        return;
    }

    if (req.url === "/health") {
        res.writeHead(200);
        res.end("OK");
        return;
    }

    if (req.url === "/metrics") {
        res.writeHead(200);
        res.end(JSON.stringify({
            memory: process.memoryUsage(),
            uptime: process.uptime()
        }));
        return;
    }

    res.writeHead(404);
    res.end("Not Found");
});

server.listen(PORT, () =>
    console.log(`Running on ${PORT}`)
);
