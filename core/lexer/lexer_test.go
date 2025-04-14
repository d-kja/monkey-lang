package lexer

import (
	. "monkey/core/tokens"
	"testing"
)

func TestLexer(t *testing.T) {
	instance := Lexer{}

	t.Run("Convert basic tokens", func(test *testing.T) {
		input := `=+(){},;[]`
		expected := []Token{
			{Type: ASSIGN, Value: "="},
			{Type: PLUS, Value: "+"},
			{Type: LEFT_PARENTESIS, Value: "("},
			{Type: RIGHT_PARENTESIS, Value: ")"},
			{Type: LEFT_BRACES, Value: "{"},
			{Type: RIGHT_BRACES, Value: "}"},
			{Type: COMMA, Value: ","},
			{Type: SEMICOLON, Value: ";"},
			{Type: LEFT_BRACKET, Value: "["},
			{Type: RIGHT_BRACKET, Value: "]"},
			{Type: EOF, Value: ""},
		}

		lexer := instance.New(input)
		for idx, exp := range expected {
			token := lexer.Read()

			if token.Value != exp.Value {
				test.Fatalf("[%d] Invalid token value, expected: %q, received: %q", idx, exp.Value, token.Value)
			}

			if token.Type != exp.Type {
				test.Fatalf("[%d] Invalid token type, expected: %q, received: %q", idx, exp.Type, token.Type)
			}
		}
	})

	t.Run("Convert identifiers", func(test *testing.T) {
		input := `let add = fn(x, y) {
  x + y;
};

let five = 5;
let ten = 10;

let result = add(five, ten);`
		expected := []Token{
			{Type: LET, Value: "let"},
			{Type: IDENTIFIER, Value: "add"},
			{Type: ASSIGN, Value: "="},
			{Type: FUNCTION, Value: "fn"},
			{Type: LEFT_PARENTESIS, Value: "("},
			{Type: IDENTIFIER, Value: "x"},
			{Type: COMMA, Value: ","},
			{Type: IDENTIFIER, Value: "y"},
			{Type: RIGHT_PARENTESIS, Value: ")"},
			{Type: LEFT_BRACES, Value: "{"},
			{Type: IDENTIFIER, Value: "x"},
			{Type: PLUS, Value: "+"},
			{Type: IDENTIFIER, Value: "y"},
			{Type: SEMICOLON, Value: ";"},
			{Type: RIGHT_BRACES, Value: "}"},
			{Type: SEMICOLON, Value: ";"},
			{Type: LET, Value: "let"},
			{Type: IDENTIFIER, Value: "five"},
			{Type: ASSIGN, Value: "="},
			{Type: INT, Value: "5"},
			{Type: SEMICOLON, Value: ";"},
			{Type: LET, Value: "let"},
			{Type: IDENTIFIER, Value: "ten"},
			{Type: ASSIGN, Value: "="},
			{Type: INT, Value: "10"},
			{Type: SEMICOLON, Value: ";"},
			{Type: LET, Value: "let"},
			{Type: IDENTIFIER, Value: "result"},
			{Type: ASSIGN, Value: "="},
			{Type: IDENTIFIER, Value: "add"},
			{Type: LEFT_PARENTESIS, Value: "("},
			{Type: IDENTIFIER, Value: "five"},
			{Type: COMMA, Value: ","},
			{Type: IDENTIFIER, Value: "ten"},
			{Type: RIGHT_PARENTESIS, Value: ")"},
			{Type: SEMICOLON, Value: ";"},
		}

		lexer := instance.New(input)
		for idx, exp := range expected {
			token := lexer.Read()

			if token.Value != exp.Value {
				test.Fatalf("[%d] Invalid token value, expected: %q, received: %q", idx, exp.Value, token.Value)
			}

			if token.Type != exp.Type {
				test.Fatalf("[%d] Invalid token type, expected: %q, received: %q", idx, exp.Type, token.Type)
			}
		}
	})

}
