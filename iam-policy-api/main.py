from fastapi import FastAPI

# Initialize the API
app = FastAPI(title="IAM Policy Generator AI", version="1.0")

# Create a simple test route (like a homepage)
@app.get("/")
def read_root():
    return {"message": "🧠 Python AI Engine is online and ready to receive code!"}

# This is the route our Go CLI will eventually send code to
@app.post("/generate")
def generate_policy(payload: dict):
    # For now, we just print what we receive to the terminal
    print(f"Received data from Go CLI: {payload}")
    
    # We will add the real Gemini AI logic here in the next step!
    return {"status": "success", "policy": "AI Policy will go here soon."}