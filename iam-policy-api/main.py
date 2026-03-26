import os
from fastapi import FastAPI
from dotenv import load_dotenv
from google import genai
from google.genai import types

# 1. Load the secret key from your .env file
load_dotenv()

# 2. Initialize the brand new Google GenAI Client
# (It automatically finds the GEMINI_API_KEY in your environment!)
try:
    client = genai.Client()
except Exception as e:
    print("⚠️ WARNING: API Key not found or invalid!", e)

# 3. Create the elite DevOps AI Persona
system_instruction = """
You are an elite AWS Cloud Security Architect. 
Your only job is to analyze infrastructure and application code (Go, Python, Terraform) 
and generate the exact, least-privilege AWS IAM JSON policy required for that code to function.
Return ONLY the raw JSON policy. Do not include markdown formatting, backticks, or any explanations.
"""

# Initialize the API
app = FastAPI(title="IAM Policy Generator AI", version="3.0")

@app.get("/")
def read_root():
    return {"message": "🧠 Python AI Engine (Powered by Gemini 2.0) is online and ready!"}

@app.post("/generate")
def generate_policy(payload: dict):
    print("📥 Received code from Go CLI. Handing off to Gemini AI...")
    
    code_context = str(payload)
    prompt = f"Analyze this codebase and generate the AWS IAM JSON policy:\n\n{code_context}"
    
    try:
        # 4. Generate the policy using the new SDK and the latest 2.0 Flash model
        response = client.models.generate_content(
            model='gemini-2.5-flash',
            contents=prompt,
            config=types.GenerateContentConfig(
                system_instruction=system_instruction,
            )
        )
        
        # Clean up the response
        clean_policy = response.text.strip().removeprefix("```json").removesuffix("```").strip()
        
        print("✅ AI successfully generated the policy! Sending it back to Go...")
        return {"status": "success", "policy": clean_policy}
        
    except Exception as e:
        print(f"❌ Error communicating with AI: {e}")
        return {"status": "error", "message": str(e)}