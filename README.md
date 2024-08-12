# llmq

`llmq` is a command-line utility that interacts with the OpenAI inference service to process piped input and provide intelligent responses. It can be used to analyze logs, manage system processes, or even answer general queries using AI models like GPT-4o-mini.

## Features

- **AI-Powered Command Line**: Send your command line output directly to an AI model for analysis and feedback.
- **Flexible Input Handling**: Works with piped input from any command-line tool.
- **Customizable**: Configure default models and API keys through a YAML configuration file.

## Installation

### Prerequisites

- **Go**: Make sure you have Go installed on your machine.
- **API Key**: You'll need an API key for the OpenAI service.

### Clone the Repository

```bash
git clone https://github.com/phildougherty/llmq.git
cd llmq
```

### Build the Project

```bash
go build
```

### Configuration

Create a configuration file at `~/.config/llmq.yaml`:

```yaml
default_model: "gpt-4o-mini"
api_key: "YOUR_API_KEY"
log_level: "info"
```

Replace `"YOUR_API_KEY"` with your actual OpenAI API key.

## Usage

### Basic Usage

You can use `llmq` by piping input into it and providing a query:

```bash
echo "toast" | ./llmq --query "What is the best topping for this?"
```

### Interesting Use Cases

Here are some creative and useful ways to integrate `llmq` into your Linux command-line environment:

#### 1. Analyze System Logs for Errors

Quickly scan logs for issues:

```bash
cat /var/log/syslog | ./llmq --query "Does anything look broken in here?"
```

#### 2. Summarize Git Commit History

Get a high-level summary of recent git commits:

```bash
git log --oneline | ./llmq --query "Summarize the recent changes in this project"
```

#### 3. Extract Key Points from a Large Text File

Summarize a large text file to get the key points:

```bash
cat large_text_file.txt | ./llmq --query "What are the key points in this text?"
```

#### 4. Suggest Improvements to a Shell Script

Get suggestions on how to improve or optimize your shell script:

```bash
cat my_script.sh | ./llmq --query "How can I improve this shell script?"
```

#### 5. Identify Resource-Intensive Processes

Find processes that are consuming a lot of resources:

```bash
ps aux | ./llmq --query "Which processes are consuming the most resources?"
```

#### 6. Generate a To-Do List from a Meeting Transcript

Convert a meeting transcript into actionable items:

```bash
cat meeting_transcript.txt | ./llmq --query "Generate a to-do list from this transcript"
```

#### 7. Summarize System Updates

Get a summary of the packages that were updated:

```bash
sudo apt-get update | ./llmq --query "Summarize the system updates"
```

### Increase Timeout

If you experience timeout issues, you can increase the timeout duration in the `ai/client.go` file:

```go
client := &http.Client{
    Timeout: 30 * time.Second,
}
```

## Supported Inference Service

Currently, `llmq` only supports OpenAI as the inference service. Ensure you have a valid API key from OpenAI to use this tool.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### Reporting Issues

If you find a bug or have a feature request, please create an issue on GitHub.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
