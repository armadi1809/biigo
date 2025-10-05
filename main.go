package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		log.Fatal("Usage: jlox [script]")
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}

}

func runFile(filepath string) {
	fileBytes, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("error reading source code file %v", err)
	}
	err = run(string(fileBytes))
	if err != nil {
		os.Exit(64)
	}

}

func runPrompt() {
	var line string
	for {
		fmt.Print("> ")
		fmt.Scan(&line)
		if len(line) == 0 {
			break
		}
		_ = run(line)
	}
}

func run(source string) error {
	return nil
}

func failure(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	log.Printf("[Line %d] Error%s: %s\n", line, where, message)
}
