import json
import numpy as np
from sklearn.datasets import fetch_california_housing
from sklearn.linear_model import LinearRegression

# Load dataset
data = fetch_california_housing()

X = data.data
y = data.target

# Train model
model = LinearRegression()
model.fit(X, y)

# Save weights
model_data = {
    "coefficients": model.coef_.tolist(),
    "intercept": model.intercept_
}

with open("model.json", "w") as f:
    json.dump(model_data, f)

print("Model saved")