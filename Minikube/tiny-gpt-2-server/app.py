import os
from fastapi import FastAPI
from pydantic import BaseModel
from transformers import AutoTokenizer, AutoModelForCausalLM
import torch

MODEL_NAME = os.getenv("MODEL_NAME", "sshleifer/tiny-gpt2")

print("Loading model:", MODEL_NAME)

tokenizer = AutoTokenizer.from_pretrained(MODEL_NAME)
model = AutoModelForCausalLM.from_pretrained(MODEL_NAME)

app = FastAPI()

class Prompt(BaseModel):
    prompt: str
    max_tokens: int = 40

@app.get("/")
def root():
    return {"msg": "Tiny Decoder API running"}

@app.get("/health")
def health():
    return {"status": "ok"}

@app.post("/generate")
def generate(p: Prompt):
    inputs = tokenizer(p.prompt, return_tensors="pt")

    outputs = model.generate(
        **inputs,
        max_length=p.max_tokens,
        do_sample=True,
        temperature=0.8
    )

    text = tokenizer.decode(outputs[0], skip_special_tokens=True)

    return {"output": text}
