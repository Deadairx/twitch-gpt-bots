package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetOpenAIResponse(message string) (string, error) {
    type Payload struct {
        Prompt string `json:"prompt"`
        MaxTokens int `json:"max_tokens"`
    }

    data := Payload{
        Prompt: message,
        MaxTokens: 50, // Adjust this to change the length of the response
    }

    payloadBytes, err := json.Marshal(data)
    if err != nil {
        return "", err
    }
    body := bytes.NewReader(payloadBytes)
    log.Println("Requesting from OpenAI:", string(payloadBytes))

    // TODO: Check API reference for the correct endpoint. I'm not sure if completions is the right one.
    req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", body)
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer " + os.Getenv("OPENAI_API_KEY"))

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    log.Println("Response from OpenAI:", resp.Status)
    defer resp.Body.Close()

    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    var result map[string]interface{}
    if err := json.Unmarshal(respBody, &result); err != nil {
        return "", err
    }

    if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
        if firstChoice, ok := choices[0].(map[string]interface{}); ok {
            if text, ok := firstChoice["text"].(string); ok {
                return text, nil
            }
        }
    }

    return "", nil
}
