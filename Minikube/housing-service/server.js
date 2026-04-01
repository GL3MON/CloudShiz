const express = require('express')
const fs = require('fs')

const app = express()
app.use(express.json())

// Load model
const model = JSON.parse(fs.readFileSync('./model.json'))

const weights = model.coefficients
const intercept = model.intercept

// prediction function
function predict(features) {

    let result = intercept

    for (let i = 0; i < weights.length; i++) {
        result += weights[i] * features[i]
    }

    return result
}

// API endpoint
app.post('/predict', (req, res) => {

    const features = req.body.features

    if (!features || features.length !== weights.length) {
        return res.status(400).send("Invalid input")
    }

    const prediction = predict(features)

    res.json({
        prediction: prediction
    })

})

app.listen(3000, () => {
    console.log("Server running on port 3000")
})