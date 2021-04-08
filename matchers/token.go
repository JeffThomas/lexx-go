package matchers

type TokenType int

const (
	SYSTEM     TokenType = iota
	WORD       TokenType = iota
	WHITESPACE TokenType = iota
	INTEGER    TokenType = iota
	FLOAT      TokenType = iota
	STRING     TokenType = iota
	SYMBOL     TokenType = iota
	KEYWORD    TokenType = iota
	IDENTIFIER TokenType = iota
)

func (mt TokenType) String() string {
	names := [...]string{
		"SYSTEM",
		"WORD",
		"WHITESPACE",
		"INTEGER",
		"SYMBOL",
		"KEYWORD"}

	if mt < INTEGER || mt > SYSTEM {
		return "Unknown"
	}

	return names[mt]
}

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}
