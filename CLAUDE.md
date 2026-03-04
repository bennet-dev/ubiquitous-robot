# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Important: Learning Mode

The user is learning Go while building this project. Do NOT implement code unless explicitly asked. Instead, explain concepts and approaches, then let the user implement the code themselves.

## Project Overview

This is a CodeCrafters challenge to build a Claude Code-like AI coding assistant in Go. The assistant uses LLMs via OpenRouter's OpenAI-compatible API to understand code and perform actions through tool calls.

## Build and Run Commands

```sh
# Build and run locally
./your_program.sh -p "your prompt here"

# Submit to CodeCrafters
codecrafters submit
```

## Environment Variables

- `OPENROUTER_API_KEY` - Required API key for OpenRouter
- `OPENROUTER_BASE_URL` - Optional, defaults to `https://openrouter.ai/api/v1`

## Architecture

Single-file Go application (`app/main.go`) using:
- `github.com/openai/openai-go/v3` - OpenAI-compatible API client
- Model: `anthropic/claude-haiku-4.5` via OpenRouter

The program takes a prompt via `-p` flag and sends it to the LLM, printing the response.
