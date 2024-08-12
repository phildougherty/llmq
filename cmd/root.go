// cmd/root.go
package cmd

import (
    "flag"
    "fmt"
    "2doc.co/llmq/ai"
    "2doc.co/llmq/config"
    "2doc.co/llmq/logger"
    "2doc.co/llmq/utils"
)

func Execute() {
    // Load config
    cfg := config.LoadConfig()

    // Setup logger
    log := logger.SetupLogger(cfg.LogLevel)

    // Parse command-line flags
    model := flag.String("model", cfg.DefaultModel, "AI model to use")
    query := flag.String("query", "", "Query to send to the AI model")
    flag.Parse()

    log.Debug("Model:", *model)
    log.Debug("Query:", *query)

    // Read from stdin
    input, err := utils.ReadStdin()
    if err != nil {
        log.Fatalf("Failed to read input: %v", err)
    }

    log.Debug("Input:", input)

    // Call AI service
    result, err := ai.CallAIService(cfg.APIKey, *model, input, *query)
    if err != nil {
        log.Fatalf("Failed to call AI service: %v", err)
    }

    log.Debug("AI Response:", result)

    // Print the result
    fmt.Println(result)
}
