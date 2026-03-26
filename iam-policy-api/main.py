import os
from fastapi import FastAPI
from dotenv import load_dotenv
import google.generativeai as genai

# 1. Load the secret key from your .env file
load_dotenv()
api_key = os.getenv("GEMINI_API_KEY")

if not api_key:
    print("⚠️ WARNING: GEMINI_API_KEY not found! Did you create the .env file?")

# 2. Configure the Google AI SDK
genai.configure(api_key=api_key)

# 3. Create the elite DevOps AI Persona
system_instruction = """
You are an elite AWS Cloud Security Architect. 
Your only job is to analyze infrastructure and application code (Go, Python, Terraform) 
and generate the exact, least-privilege AWS IAM JSON policy required for that code to function.
Return ONLY the raw JSON policy. Do not include markdown formatting, backticks, or any explanations.
"""

# Initialize the blazing fast Gemini 1.5 Flash model
model = genai.GenerativeModel(
    model_name="gemini-1.5-flash",
    system_instruction=system_instruction
)

# Initialize the API
app = FastAPI(title="IAM Policy Generator AI", version="2.0")

@app.get("/")
def read_root():
    return {"message": "🧠 Python AI Engine (Powered by Gemini) is online and ready!"}

@app.post("/generate")
def generate_policy(payload: dict):
    print("📥 Received code from Go CLI. Handing off to Gemini AI...")
    
    # Convert the incoming JSON payload into a massive string for the AI to read
    code_context = str(payload)
    
    # Build the final prompt
    prompt = f"Analyze this codebase and generate the AWS IAM JSON policy:\n\n{code_context}"
    
    try:
        # 4. Generate the policy using the AI
        response = model.generate_content(prompt)
        
        # Clean up the response just in case the AI stubbornly includes markdown formatting
        clean_policy = response.text.strip().removeprefix("```json").removesuffix("```").strip()
        
        print("✅ AI successfully generated the policy! Sending it back to Go...")
        return {"status": "success", "policy": clean_policy}
        
    except Exception as e:
        print(f"❌ Error communicating with AI: {e}")
        return {"status": "error", "message": str(e)}