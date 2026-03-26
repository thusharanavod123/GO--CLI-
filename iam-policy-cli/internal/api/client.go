package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"iam-policy-cli/internal/scanner"
)

type Payload struct {
	Files []scanner.FileData `json:"files"`
}

// APIResponse matches the JSON structure our Python server sends back
type APIResponse struct {
	Status  string `json:"status"`
	Policy  string `json:"policy"`
	Message string `json:"message"`
}

func SendToAI(files []scanner.FileData) error {
	fmt.Println("\n🌐 Connecting to Python AI Engine at http://127.0.0.1:8000...")

	payload := Payload{Files: files}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to package data: %w", err)
	}

	apiUrl := "http://127.0.0.1:8000/generate"
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to connect to API: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	// --- NEW FORMATTING LOGIC ---
	// 1. Unmarshal the outer response from Python
	var apiResp APIResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return fmt.Errorf("failed to parse API response: %w", err)
	}

	// Handle errors from the Python side
	if apiResp.Status == "error" {
		return fmt.Errorf("API Error: %s", apiResp.Message)
	}

	// 2. Take the inner policy string and format it beautifully
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, []byte(apiResp.Policy), "", "  ")
	if err != nil {
		// If it fails to indent (e.g., AI sent bad JSON), just print the raw string
		fmt.Println("\n✅ AI Engine Responded:\n")
		fmt.Println(apiResp.Policy)
		return nil
	}

	// Print the beautiful JSON!
	fmt.Println("\n✅ Successfully Generated IAM Policy!")
	fmt.Println("--------------------------------------------------")
	fmt.Println(prettyJSON.String())
	fmt.Println("--------------------------------------------------")

	return nil
}