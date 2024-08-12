// ai/client.go
package ai

import (
    "bytes"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "time"
    "2doc.co/llmq/logger"
)

// AIResponse represents the structure of the response from the AI API
type AIResponse struct {
    Choices []struct {
        Message struct {
            Role    string `json:"role"`
            Content string `json:"content"`
        } `json:"message"`
    } `json:"choices"`
}

func CallAIService(apiKey, model, input, query string) (string, error) {
    log := logger.SetupLogger("debug")

    apiUrl := "https://api.openai.com/v1/chat/completions"

    body := map[string]interface{}{
        "model": model,
        "messages": []map[string]string{
            {"role": "user", "content": query + ": " + input},
        },
    }

    jsonData, err := json.Marshal(body)
    if err != nil {
        log.Errorf("Error marshalling request body: %v", err)
        return "", err
    }

    req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(jsonData))
    if err != nil {
        log.Errorf("Error creating request: %v", err)
        return "", err
    }

    req.Header.Set("Authorization", "Bearer "+apiKey)
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{
        Timeout: 30 * time.Second,  // Increased timeout to 30 seconds
    }

    resp, err := client.Do(req)
    if err != nil {
        log.Errorf("Error making API request: %v", err)
        return "", err
    }
    defer resp.Body.Close()

    bodyBytes, _ := ioutil.ReadAll(resp.Body)
    log.Debugf("Raw API Response: %s", string(bodyBytes))

    if resp.StatusCode != http.StatusOK {
        log.Errorf("API request failed with status: %v", resp.Status)
        return "", errors.New("API request failed")
    }

    var aiResp AIResponse
    err = json.Unmarshal(bodyBytes, &aiResp)
    if err != nil {
        log.Errorf("Error decoding response: %v", err)
        return "", err
    }

    if len(aiResp.Choices) == 0 {
        log.Warn("No choices in API response")
        return "", nil
    }

    // Return the content of the first choice
    return aiResp.Choices[0].Message.Content, nil
}
