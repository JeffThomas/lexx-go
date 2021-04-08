package matchers

type MatcherResult struct {
	Token      *Token
	Err        error
	Precedence int8
}
