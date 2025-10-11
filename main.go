package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/biigo/interpreter"
	"github.com/armadi1809/biigo/lexer"
	"github.com/armadi1809/biigo/parser"
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
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		err := run(line)
		if err != nil {
			log.Print(err.Error())
		}
	}
}

func run(source string) error {
	// lexing
	lex := lexer.NewLexer(source)
	toks, err := lex.ScanTokens()

	if err != nil {
		return err
	}

	// parsing
	parser := parser.NewParser(toks)
	exp, err := parser.Parse()

	if err != nil {
		return err
	}

	val, err := interpreter.Interpret(exp)
	if err == nil {
		fmt.Printf("%v\n", val)
	}
	return err
}
