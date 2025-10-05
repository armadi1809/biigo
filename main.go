package main

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/biigo/lexer"
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
		err := run(line)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func run(source string) error {
	lex := lexer.NewLexer(source)
	toks, err := lex.ScanTokens()

	if err != nil {
		return err
	}

	fmt.Println(toks)
	return nil
}
