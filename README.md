# Go Coding Agent

A conversational AI coding agent implemented in Go that leverages large language models to autonomously read files, write files, and execute shell commands based on user prompts.

## Overview

This application is an agentic system that uses Claude (via OpenRouter) to interpret natural language prompts and take actions on your behalf. The agent has access to three tools:

- **Read**: Read and return the contents of a file
- **Write**: Write content to a file
- **Bash**: Execute shell commands

The agent processes your initial prompt, makes decisions about which tools to use, executes them, and incorporates the results back into the conversation until it completes the task or provides a final response.

## Features

- **Autonomous Tool Use**: The LLM can decide which tools to use and in what sequence
- **Agentic Loop**: Continuously processes responses and tool results until the task is complete
- **File Operations**: Read and write files on your filesystem
- **Command Execution**: Execute arbitrary shell commands
- **OpenRouter Integration**: Uses Claude Haiku via OpenRouter API for efficient and cost-effective processing

## Prerequisites

- Go 1.20 or higher
- OpenRouter API key (get one at https://openrouter.ai)

## Installation

```bash
go build -o coding-agent ./app
```

## Configuration

Set the following environment variables:

```bash
export OPENROUTER_API_KEY="your-api-key-here"
# Optional: defaults to https://openrouter.ai/api/v1 if not set
export OPENROUTER_BASE_URL="https://openrouter.ai/api/v1"
```

## Usage

Run the agent with a prompt describing the task:

```bash
./coding-agent -p "Your task description here"
```

### Examples

Read a file:
```bash
./coding-agent -p "Read the contents of README.md and tell me what this project does"
```

Create a new file:
```bash
./coding-agent -p "Create a new file called hello.txt with the content 'Hello, World!'"
```

Execute commands:
```bash
./coding-agent -p "List all files in the current directory"
```

Complex tasks:
```bash
./coding-agent -p "Create a Python script that prints the first 10 Fibonacci numbers and save it as fibonacci.py, then run it"
```

## How It Works

1. **Initialization**: The application starts with your prompt sent to the Claude Haiku model via OpenRouter
2. **Agent Loop**: 
   - The model receives your prompt and the list of available tools
   - Claude decides which tools to use and calls them with appropriate arguments
   - Results are returned to Claude in the conversation history
   - This process repeats until Claude determines the task is complete
3. **Output**: The final response from Claude is printed to stdout
4. **Logging**: Tool calls and arguments are logged to stderr for debugging

## Project Structure

```
app/
├── main.go      # Entry point and main agent loop
├── client.go    # OpenRouter API client setup and tool definitions
├── tools.go     # Implementation of Read, Write, and Bash tools
└── utils.go     # Utility functions
```

## Architecture

The application follows a tool-use pattern where:

1. Tools are defined as OpenAI-compatible function definitions
2. The LLM receives these tool definitions and can choose to call them
3. Tool results are fed back into the conversation context
4. The LLM continues processing until it has completed the task or provided a response

## Safety Considerations

⚠️ **Warning**: This agent can execute arbitrary shell commands and modify files on your system. Only run with prompts you trust, as the LLM may interpret requests in unexpected ways.

## Limitations

- The agent uses Claude Haiku for speed and cost-efficiency, which may have limitations on complex reasoning tasks
- Single-threaded execution (processes one prompt at a time)
- No persistent state between runs

## Dependencies

- `github.com/openai/openai-go/v3` - OpenAI-compatible Go client for OpenRouter

## License

MIT
