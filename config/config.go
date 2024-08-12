// config/config.go
package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
)

type Config struct {
    DefaultModel string `yaml:"default_model"`
    APIKey       string `yaml:"api_key"`
    LogLevel     string `yaml:"log_level"`
}

func LoadConfig() Config {
    configPath := os.Getenv("HOME") + "/.config/llmq.yaml"
    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        log.Fatalf("Failed to read config file: %v", err)
    }

    var cfg Config
    err = yaml.Unmarshal(data, &cfg)
    if err != nil {
        log.Fatalf("Failed to parse config file: %v", err)
    }

    return cfg
}
