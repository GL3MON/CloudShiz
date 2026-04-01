const express = require("express")
const app = express()

app.get("/", (req, res) => {
    res.send("Hiii")
})

app.listen(3000)