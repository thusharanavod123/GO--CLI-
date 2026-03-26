package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"iam-policy-cli/internal/scanner"
)

// Payload represents the exact JSON structure we are sending to Python
type Payload struct {
	Files []scanner.FileData `json:"files"`
}

// SendToAI shoots the scooped-up files over to your Python FastAPI backend
func SendToAI(files []scanner.FileData) error {
	fmt.Println("\n🌐 Connecting to Python AI Engine at http://127.0.0.1:8000...")

	// 1. Package the files into our JSON format
	payload := Payload{Files: files}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to package data: %w", err)
	}

	// 2. Fire the data over the internet to the Python server
	apiUrl := "http://127.0.0.1:8000/generate"
	resp, err := http.Post(apiUrl, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to connect to API (is the Python server running?): %w", err)
	}
	// Always close the connection when we are done
	defer resp.Body.Close()

	// 3. Read the Python AI's response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	fmt.Println("✅ AI Engine Responded!")
	fmt.Println("--------------------------------------------------")
	fmt.Println(string(body))
	fmt.Println("--------------------------------------------------")

	return nil
}