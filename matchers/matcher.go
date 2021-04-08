package matchers

import (
	token_class "github.com/JeffThomas/lexx/token"
)

type MatcherResult struct {
	Token      *token_class.Token
	Err        error
	Precedence int8
}
