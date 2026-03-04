package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

func main() {
	var messages []openai.ChatCompletionMessageParamUnion
	var prompt string
	flag.StringVar(&prompt, "p", "", "Prompt to send to LLM")
	flag.Parse()

	if prompt == "" {
		panic("Prompt must not be empty")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	baseUrl := os.Getenv("OPENROUTER_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://openrouter.ai/api/v1"
	}

	if apiKey == "" {
		panic("Env variable OPENROUTER_API_KEY not found")
	}

	messages = append(messages, openai.ChatCompletionMessageParamUnion{
		OfUser: &openai.ChatCompletionUserMessageParam{
			Content: openai.ChatCompletionUserMessageParamContentUnion{
				OfString: openai.String(prompt),
			},
		},
	})

	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL(baseUrl))

	for {
		resp, err := createChatCompletion(&client, messages)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}

		// Translate assistant response into param type and append
		toolCalls := resp.Choices[0].Message.ToolCalls
		messages = append(messages, assistantMsgToParam(resp.Choices[0].Message))

		// No tool calls means the model is done — print final response and exit
		if len(toolCalls) == 0 {
			fmt.Print(resp.Choices[0].Message.Content)
			break
		}

		// Execute tool calls and append results
		for _, toolCall := range toolCalls {
			switch toolCall.Function.Name {
			case "Read":
				args := toolCall.Function.Arguments

				fmt.Fprintf(os.Stderr, "Read tool called with arguments: %v\n", args)
				content, err := read(args)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
				// Append tool result to messages
				messages = append(messages, openai.ChatCompletionMessageParamUnion{
					OfTool: &openai.ChatCompletionToolMessageParam{
						ToolCallID: toolCall.ID,
						Content: openai.ChatCompletionToolMessageParamContentUnion{
							OfString: openai.String(string(content)),
						},
					},
				})
			case "Write":
				args := toolCall.Function.Arguments

				fmt.Fprintf(os.Stderr, "Write tool called with arguments: %v\n", args)
				if err := write(args); err != nil {
					fmt.Fprintf(os.Stderr, "error: %v\n", err)
					os.Exit(1)
				}
				// Append success message to messages
				messages = append(messages, openai.ChatCompletionMessageParamUnion{
					OfTool: &openai.ChatCompletionToolMessageParam{
						ToolCallID: toolCall.ID,
						Content: openai.ChatCompletionToolMessageParamContentUnion{
							OfString: openai.String("Write successful"),
						},
					},
				})
				// handle write_file
			default:
				// unknown tool
			}
		}
		// Loop back — calls API again with the full conversation history
	}
}
