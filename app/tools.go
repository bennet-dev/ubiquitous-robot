package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func read(args string) (string, error) {
	var readArgs ReadFileArgs
	if err := json.Unmarshal([]byte(args), &readArgs); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	content, err := os.ReadFile(readArgs.Path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

type ReadFileArgs struct {
	Path string `json:"file_path"`
}
