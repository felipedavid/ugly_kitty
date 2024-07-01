package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/felipedavid/ugly_kitty/lexer"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Fprintf(os.Stderr, "Usage: ./uk <script>")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		err := runFile(os.Args[1])
		if err != nil {
			panic(err)
		}
	} else {
		prompt()
	}
}

func prompt() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "quit" {
			break
		}

		if line == "" {
			fmt.Printf("> ")
			continue
		}

		err := run([]byte(line))
		if err != nil {
			fmt.Printf("[ERROR] %s", err.Error())
		}

		fmt.Printf("> ")
	}
}

func runFile(fileName string) error {
	source, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	return run(source)
}

func run(source []byte) error {
	lex := lexer.New(source)

	tokens := lex.ParseTokens()
	for _, token := range tokens {
		fmt.Printf("%s\n", token.String())
	}

	return nil
}
