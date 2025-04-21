package tokens

type _Type = string

type Token struct {
	Type  _Type
	Value string
}

const (
	EOF        = "EOF"
	ILLEGAL    = "ILLEGAL"
	IDENTIFIER = "IDENTIFIER"

	// TYPES
	INT     = "INT"
	FLOAT   = "FLOAT"
	STRING  = "STRING"
	BOOLEAN = "BOOLEAN"

	// OPERATIONS
	PLUS              = "+"
	MINUS             = "-"
	DIVISION          = "/"
	MULTIPLICATION    = "*"
	NOT               = "!"
	SMALLER_THAN      = "<"
	GREATER_THAN      = ">"
	SMALLER_OR_EQUALS = "<="
	GREATER_OR_EQUALS = ">="
	EQUALS            = "=="
	DIFFERENT         = "!="

	// SPECIAL
	COMMA            = ","
	ASSIGN           = "="
	SEMICOLON        = ";"
	LEFT_BRACES      = "{"
	RIGHT_BRACES     = "}"
	LEFT_BRACKET     = "["
	RIGHT_BRACKET    = "]"
	LEFT_PARENTESIS  = "("
	RIGHT_PARENTESIS = ")"

	// KEYWORDS
	LET      = "LET"
	RETURN   = "RETURN"
	FUNCTION = "FUNCTION"
	ELSE     = "ELSE"
	IF       = "IF"
)

var identifiers = map[string]_Type{
	"let":    LET,
	"fn":     FUNCTION,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   BOOLEAN,
	"false":  BOOLEAN,
}

func (self Token) LookupIdentifier() _Type {
	if value, ok := identifiers[self.Value]; ok {
		return value
	}

	return IDENTIFIER
}

func NewToken(_type _Type, input byte) Token {
	return Token{
		Type:  _type,
		Value: string(input),
	}
}

func NewTokenAsString(_type _Type, input string) Token {
	return Token{
		Type:  _type,
		Value: input,
	}
}
