package lexer

import (
	"monkey/core/tokens"
	"strings"
)

type Lexer struct {
	Char         byte
	Input        string
	Position     uint
	NextPosition uint
}

func (self Lexer) New(input string) *Lexer {
	instance := Lexer{
		Input: input,
	}

	instance.Read()
	return &instance
}

func (self Lexer) Peek() byte {
	if self.NextPosition >= uint(len(self.Input)) {
		return 0
	}

	return self.Input[self.NextPosition]
}

func (self *Lexer) readChar() {
	if self.NextPosition >= uint(len(self.Input)) {
		self.Char = 0
	} else {
		self.Char = self.Input[self.NextPosition]
	}

	self.Position = self.NextPosition
	self.NextPosition++
}

func (self *Lexer) readIdentifier() string {
	position := self.Position

	for self.isLetter(self.Char) {
		self.readChar()
	}

	return self.Input[position:self.Position]
}

func (self *Lexer) readNumber() tokens.Token {
	token := tokens.Token{}
	position := self.Position

	for self.isNumber(self.Char) {
		self.readChar()
	}

	token.Value = self.Input[position:self.Position]
	token.Type = tokens.INT

	if strings.Contains(token.Value, ".") {
		token.Type = tokens.FLOAT
	}

	return token
}

func (self *Lexer) eatWhitespaces() {
	for self.isWhitespace(self.Char) {
		self.readChar()
	}
}

func (self *Lexer) isLetter(char byte) bool {
	if 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_' {
		return true
	}

	return false
}

func (self *Lexer) isNumber(char byte) bool {
	if '0' <= char && char <= '9' || char == '.' {
		return true
	}

	return false
}

func (self *Lexer) isWhitespace(char byte) bool {
	if char == ' ' || char == '\n' || char == '\r' || char == '\t' {
		return true
	}

	return false
}

func (self *Lexer) Read() tokens.Token {
	var token tokens.Token
	self.eatWhitespaces()

	switch {
	case self.Char == '=':
		{
			nextToken := self.Peek()

			if nextToken == '=' {
				token = tokens.NewTokenAsString(tokens.EQUALS, string(self.Char) + string(nextToken))
				self.readChar()

				break
			}

			token = tokens.NewToken(tokens.ASSIGN, self.Char)
		}
	case self.Char == '!':
		{
			nextToken := self.Peek()

			if nextToken == '=' {
				token = tokens.NewTokenAsString(tokens.DIFFERENT, string(self.Char) + string(nextToken))
				self.readChar()

				break
			}

			token = tokens.NewToken(tokens.NOT, self.Char)
		}
	case self.Char == '<':
		{
			nextToken := self.Peek()

			if nextToken == '=' {
				token = tokens.NewTokenAsString(tokens.SMALLER_OR_EQUALS, string(self.Char) + string(nextToken))
				self.readChar()

				break
			}

			token = tokens.NewToken(tokens.SMALLER_THAN, self.Char)
		}
	case self.Char == '>':
		{
			nextToken := self.Peek()

			if nextToken == '=' {
				token = tokens.NewTokenAsString(tokens.GREATER_OR_EQUALS, string(self.Char) + string(nextToken))
				self.readChar()

				break
			}

			token = tokens.NewToken(tokens.GREATER_THAN, self.Char)
		}
	case self.Char == '+':
		{
			token = tokens.NewToken(tokens.PLUS, self.Char)
		}
	case self.Char == '-':
		{
			token = tokens.NewToken(tokens.MINUS, self.Char)
		}
	case self.Char == '*':
		{
			token = tokens.NewToken(tokens.MULTIPLICATION, self.Char)
		}
	case self.Char == '/':
		{
			token = tokens.NewToken(tokens.DIVISION, self.Char)
		}
	case self.Char == ',':
		{
			token = tokens.NewToken(tokens.COMMA, self.Char)
		}
	case self.Char == ';':
		{
			token = tokens.NewToken(tokens.SEMICOLON, self.Char)
		}
	case self.Char == '[':
		{
			token = tokens.NewToken(tokens.LEFT_BRACKET, self.Char)
		}
	case self.Char == ']':
		{
			token = tokens.NewToken(tokens.RIGHT_BRACKET, self.Char)
		}
	case self.Char == '(':
		{
			token = tokens.NewToken(tokens.LEFT_PARENTESIS, self.Char)
		}
	case self.Char == ')':
		{
			token = tokens.NewToken(tokens.RIGHT_PARENTESIS, self.Char)
		}
	case self.Char == '{':
		{
			token = tokens.NewToken(tokens.LEFT_BRACES, self.Char)
		}
	case self.Char == '}':
		{
			token = tokens.NewToken(tokens.RIGHT_BRACES, self.Char)
		}
	case self.Char == 0:
		{
			token.Type = tokens.EOF
			token.Value = ""
		}

	default:
		{
			switch {
			case self.isLetter(self.Char):
				{
					token.Value = self.readIdentifier()
					token.Type = token.LookupIdentifier()

					return token
				}
			case self.isNumber(self.Char):
				{
					token = self.readNumber()

					return token
				}

			default:
				{
					token = tokens.NewToken(tokens.ILLEGAL, self.Char)
				}
			}
		}
	}

	self.readChar()
	return token
}
