package main

import (
	"context"
	"fmt"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/shared"
)

var tools = []openai.ChatCompletionToolUnionParam{
	openai.ChatCompletionFunctionTool(shared.FunctionDefinitionParam{
		Name:        "Read",
		Description: openai.Opt("Read and return the contents of a file"),
		Parameters: shared.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"file_path": map[string]any{
					"type":        "string",
					"description": "The path to the file to read",
				},
			},
			"required": []string{"file_path"},
		},
	}),
	openai.ChatCompletionFunctionTool(shared.FunctionDefinitionParam{
		Name:        "Write",
		Description: openai.Opt("Write content to a file"),
		Parameters: shared.FunctionParameters{
			"type": "object",
			"properties": map[string]any{
				"file_path": map[string]any{
					"type":        "string",
					"description": "The path to the file to write to",
				},
				"content": map[string]any{
					"type":        "string",
					"description": "The content to write to the file",
				},
			},
			"required": []string{"file_path", "content"},
		},
	}),
	openai.ChatCompletionFunctionTool(shared.FunctionDefinitionParam{
		Name:        "Bash",
		Description: openai.Opt("Execute a shell command"),
		Parameters: shared.FunctionParameters{
			"type":     "object",
			"required": []string{"command"},
			"properties": map[string]any{
				"command": map[string]any{
					"type":        "string",
					"description": "The command to execute",
				},
			},
		},
	}),
}

func createChatCompletion(client *openai.Client, messages []openai.ChatCompletionMessageParamUnion) (*openai.ChatCompletion, error) {
	resp, err := client.Chat.Completions.New(context.Background(),
		openai.ChatCompletionNewParams{
			Model:    "anthropic/claude-haiku-4.5",
			Messages: messages,
			Tools:    tools,
		},
	)
	if err != nil {
		return nil, err
	}
	if len(resp.Choices) == 0 {
		return nil, fmt.Errorf("no choices in response")
	}
	return resp, nil
}

func assistantMsgToParam(msg openai.ChatCompletionMessage) openai.ChatCompletionMessageParamUnion {
	var toolCallParams []openai.ChatCompletionMessageToolCallUnionParam
	for _, tc := range msg.ToolCalls {
		toolCallParams = append(toolCallParams, openai.ChatCompletionMessageToolCallUnionParam{
			OfFunction: &openai.ChatCompletionMessageFunctionToolCallParam{
				ID: tc.ID,
				Function: openai.ChatCompletionMessageFunctionToolCallFunctionParam{
					Name:      tc.Function.Name,
					Arguments: tc.Function.Arguments,
				},
			},
		})
	}

	return openai.ChatCompletionMessageParamUnion{
		OfAssistant: &openai.ChatCompletionAssistantMessageParam{
			Content: openai.ChatCompletionAssistantMessageParamContentUnion{
				OfString: openai.String(msg.Content),
			},
			ToolCalls: toolCallParams,
		},
	}
}
