package core

import (
	"bufio"
	"fmt"
	"io"
	"monkey/core/lexer"
	"monkey/core/tokens"
)

type Repl struct{}

func (self *Repl) Run(input io.Reader, output io.Writer) {
	reader := bufio.NewScanner(input)
	instance := lexer.Lexer{}

	for {
		fmt.Printf("> ")
		hasInput := reader.Scan()
		if !hasInput {
			continue
		}

		text := reader.Text()
		if text == "exit" {
			break
		}

		lexer := instance.New(text)

		for token := lexer.Read(); token.Type != tokens.EOF; token = lexer.Read() {
			fmt.Printf("[REPL] Input: %+v, token: %+v\n", token.Value, token.Type)
		}
	}
}
