package matchers

type TokenType int

const (
	SYSTEM      TokenType = iota
	WORD        TokenType = iota
	WHITESPACE  TokenType = iota
	INTEGER     TokenType = iota
	FLOAT       TokenType = iota
	STRING      TokenType = iota
	OPERATOR    TokenType = iota
	KEYWORD     TokenType = iota
	IDENTIFIER  TokenType = iota
	SYMBOLIC    TokenType = iota
	PUNCTUATION TokenType = iota
	UNDEFINED   TokenType = iota
	CUSTOM      TokenType = iota
)

func (mt TokenType) String() string {
	names := [...]string{
		"SYSTEM",
		"WORD",
		"WHITESPACE",
		"INTEGER",
		"FLOAT",
		"STRING",
		"OPERATOR",
		"KEYWORD",
		"IDENTIFIER",
		"SYMBOLIC",
		"PUNCTUATION",
		"UNDEFINED",
		"CUSTOM",
	}

	if mt < SYSTEM || mt > UNDEFINED {
		return "Unknown"
	}

	return names[mt]
}

func (mt TokenType) MarshalText() (text []byte, err error) {
	sv := mt.String()
	return []byte(sv), nil
}

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

func (t *Token) Equals(ot *Token) bool {
	return t.Value == ot.Value && t.Type == ot.Type
}
