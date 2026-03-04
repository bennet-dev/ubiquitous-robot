package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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

func write(args string) error {
	var writeArgs WriteFileArgs
	if err := json.Unmarshal([]byte(args), &writeArgs); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return os.WriteFile(writeArgs.Path, []byte(writeArgs.Content), 0644)
}

type WriteFileArgs struct {
	Path    string `json:"file_path"`
	Content string `json:"content"`
}

func bash(args string) (string, error) {
	var bashArgs BashArgs
	if err := json.Unmarshal([]byte(args), &bashArgs); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	cmd := exec.Command("bash", "-c", bashArgs.Command)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

type BashArgs struct {
	Command string `json:"command"`
}
