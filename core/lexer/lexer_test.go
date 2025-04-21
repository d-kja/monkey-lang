package lexer

import (
	. "monkey/core/tokens"
	"testing"
)

// YES. it's incrementing, I could have removed the duplicated code, but I'll leave it here in case I want to revisit to see the whole progress

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

	t.Run("Test operator tokens", func(test *testing.T) {
		input := `
			let five = 5;
			let ten = 10;

			let add = fn(x, y) {
				x + y;
			};

			let result = add(five, ten);

			!-/*5;

			5 < 10 > 5;

			if (5 < 10) {
				return true;
			} else {
				return false;
			}

			10 == 10;
			10 != 9;
		`
		expected := []Token{
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
			{Type: IDENTIFIER, Value: "result"},
			{Type: ASSIGN, Value: "="},
			{Type: IDENTIFIER, Value: "add"},
			{Type: LEFT_PARENTESIS, Value: "("},
			{Type: IDENTIFIER, Value: "five"},
			{Type: COMMA, Value: ","},
			{Type: IDENTIFIER, Value: "ten"},
			{Type: RIGHT_PARENTESIS, Value: ")"},
			{Type: SEMICOLON, Value: ";"},
			{Type: NOT, Value: "!"},
			{Type: MINUS, Value: "-"},
			{Type: DIVISION, Value: "/"},
			{Type: MULTIPLICATION, Value: "*"},
			{Type: INT, Value: "5"},
			{Type: SEMICOLON, Value: ";"},
			{Type: INT, Value: "5"},
			{Type: SMALLER_THAN, Value: "<"},
			{Type: INT, Value: "10"},
			{Type: GREATER_THAN, Value: ">"},
			{Type: INT, Value: "5"},
			{Type: SEMICOLON, Value: ";"},
			{Type: IF, Value: "if"},
			{Type: LEFT_PARENTESIS, Value: "("},
			{Type: INT, Value: "5"},
			{Type: SMALLER_THAN, Value: "<"},
			{Type: INT, Value: "10"},
			{Type: RIGHT_PARENTESIS, Value: ")"},
			{Type: LEFT_BRACES, Value: "{"},
			{Type: RETURN, Value: "return"},
			{Type: BOOLEAN, Value: "true"},
			{Type: SEMICOLON, Value: ";"},
			{Type: RIGHT_BRACES, Value: "}"},
			{Type: ELSE, Value: "else"},
			{Type: LEFT_BRACES, Value: "{"},
			{Type: RETURN, Value: "return"},
			{Type: BOOLEAN, Value: "false"},
			{Type: SEMICOLON, Value: ";"},
			{Type: RIGHT_BRACES, Value: "}"},
			{Type: INT, Value: "10"},
			{Type: EQUALS, Value: "=="},
			{Type: INT, Value: "10"},
			{Type: SEMICOLON, Value: ";"},
			{Type: INT, Value: "10"},
			{Type: DIFFERENT, Value: "!="},
			{Type: INT, Value: "9"},
			{Type: SEMICOLON, Value: ";"},
		}

		lexer := instance.New(input)
		for idx, item := range expected {
			token := lexer.Read()

			println(idx, item.Type, item.Value, token.Type, token.Value)
			if token.Value != item.Value {
				test.Fatalf("[%d] Invalid token value, expected: %q, received: %q", idx, item.Value, token.Value)
			}

			if token.Type != item.Type {
				test.Fatalf("[%d] Invalid token type, expected: %q, received: %q", idx, item.Type, token.Type)
			}
		}
	})
}
