package lexx

import (
	"bufio"
	"fmt"
	"github.com/JeffThomas/lexx/matchers"
	"log"
	"strings"
	"testing"
)

func TestLexxNoMatchers(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("adfasdfasf"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	to, err := lexxx.GetNextToken()

	if to.Type != matchers.UNDEFINED {
		t.Errorf("Token type should be UNDEFINED not %s\n", to.Type)
	}
	if to.Value != "a" {
		t.Errorf("Token value should be 'a' not '%s'\n", err.Error())
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}

	to, err = lexxx.GetNextToken()

	if to.Type != matchers.UNDEFINED {
		t.Errorf("Token type should be UNDEFINED not %s\n", to.Type)
	}
	if to.Value != "d" {
		t.Errorf("Token value should be 'd' not '%s'\n", err.Error())
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxNullString(t *testing.T) {
	r := bufio.NewReader(strings.NewReader(""))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	lexxx.AddMatcher(matchers.StartWordMatcher)
	to, err := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("LexxToken should not be empty on EOF.\n")
	} else if to.Type != matchers.SYSTEM {
		t.Errorf("EOF tokens should be of type  SYSTEM: %s.\n", to.Type)
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func TestLexxFindsAToken(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("Text"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	lexxx.AddMatcher(matchers.StartWordMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
}

func TestLexxReFindsATokenWhenPushedBack(t *testing.T) {
	// used by precedence parsing
	// a quick and dirty 1 token rewind
	r := bufio.NewReader(strings.NewReader("Text"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	lexxx.AddMatcher(matchers.StartWordMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
	lexxx.PushToken()
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
}

func TestLexxLineCountToken(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("\n \n\n \n "))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	lexxx.AddMatcher(matchers.StartWhitespaceMatcher)
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "\n \n\n \n " {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found word tokens should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (lexxx.State.LineNext != 4 || lexxx.State.ColumnNext != 1) {
		t.Errorf("Line should be 4 and Column should be 1. Found: %d, %d", to.Line, to.Column)
	}
}

func TestLexxParsesMultiples(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("Text this	\nthing"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.SetInput(r)
	lexxx.AddMatchers([]matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.StartWhitespaceMatcher,
	})
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "Text" {
		t.Errorf("Found word should be 'Text' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find first Whitespace\n")
	} else if to.Value != " " {
		t.Errorf("Found word should be ' ' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found whitespace tokens should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 4) {
		t.Errorf("Line should be 0 and Column should be 4. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "this" {
		t.Errorf("Found word should be 'this' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 5) {
		t.Errorf("Line should be 0 and Column should be 5. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find second Whitespace\n")
	} else if to.Value != "	\n" {
		t.Errorf("Found word should be '	\n' not %s\n", to.Value)
	} else if to.Type != matchers.WHITESPACE {
		t.Errorf("Found whitespace tokens should be type WHITESPACE not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 9) {
		t.Errorf("Line should be 0 and Column should be 9. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Word\n")
	} else if to.Value != "thing" {
		t.Errorf("Found word should be 'thing' not %s\n", to.Value)
	} else if to.Type != matchers.WORD {
		t.Errorf("Found word tokens should be type WORD not %s\n", to.Type)
	}
	if to != nil && (to.Line != 1 || to.Column != 0) {
		t.Errorf("Line should be 1 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
}

func TestLexxRewindNothingToReqind(t *testing.T) {
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.SetStringInput("NoInput")
	lexxx.AddMatcher(matchers.ConfigOperatorMatcher([]string{"!++"}))
	lexxx.Rewind()
	if len(lexxx.cache) > 0 {
		t.Error("Rewind did something very strange.")
	}
}

func TestLexxPrecedence(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("newThing"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.input = r
	lexxx.AddMatcher(matchers.ConfigKeywordMatcher([]string{"new"}))
	lexxx.AddMatcher(matchers.StartWordMatcher)
	lexxx.AddMatcher(matchers.StartIdentifierMatcher)
	to, _ := lexxx.GetNextToken()

	if to == nil {
		t.Errorf("Lexx did not find Identifier\n")
	} else if to.Value != "newThing" {
		t.Errorf("Found identifier should be 'newThing' not %s\n", to.Value)
	} else if to.Type != matchers.IDENTIFIER {
		t.Errorf("Found tokens should be type IDENTIFIER not %s\n", to.Type)
	}
}

func TestLexxRewind(t *testing.T) {
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.SetStringInput("!++=+")
	lexxx.AddMatcher(matchers.ConfigOperatorMatcher([]string{"!++"}))
	to, _ := lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!++" {
		t.Errorf("Found word should be '!++' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found word tokens should be type OPERATOR not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	lexxx.Rewind()
	lexxx.Matchers = lexxx.Matchers[0:0]
	lexxx.AddMatcher(matchers.ConfigOperatorMatcher([]string{"+", "!", "+="}))
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "!" {
		t.Errorf("Found word should be '!' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found word tokens should be type OPERATOR not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 0) {
		t.Errorf("Line should be 0 and Column should be 0. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found word tokens should be type OPERATOR not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 1) {
		t.Errorf("Line should be 0 and Column should be 1. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+=" {
		t.Errorf("Found word should be '+=' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found word tokens should be type OPERATOR not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 3) {
		t.Errorf("Line should be 0 and Column should be 3. Found: %d, %d", to.Line, to.Column)
	}
	to, _ = lexxx.GetNextToken()
	if to == nil {
		t.Errorf("Lexx did not find Symbols\n")
	} else if to.Value != "+" {
		t.Errorf("Found word should be '+' not %s\n", to.Value)
	} else if to.Type != matchers.OPERATOR {
		t.Errorf("Found word tokens should be type OPERATOR not %s\n", to.Type)
	}
	if to != nil && (to.Line != 0 || to.Column != 5) {
		t.Errorf("Line should be 0 and Column should be 5. Found: %d, %d", to.Line, to.Column)
	}
}

func TestLexxSkippingOverUndefined(t *testing.T) {
	r := bufio.NewReader(strings.NewReader("+a-"))
	lexxx := Lexx{State: &StateLexx{}}
	lexxx.AddMatcher(matchers.ConfigOperatorMatcher([]string{"+", "-"}))
	lexxx.input = r
	to, err := lexxx.GetNextToken()

	if to.Type != matchers.OPERATOR {
		t.Errorf("Token type should be OPERATOR not %s\n", to.Type)
	}
	if to.Value != "+" {
		t.Errorf("Token value should be '+' not '%s'\n", err.Error())
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}

	to, err = lexxx.GetNextToken()

	if to.Type != matchers.UNDEFINED {
		t.Errorf("Token type should be UNDEFINED not %s\n", to.Type)
	}
	if to.Value != "a" {
		t.Errorf("Token value should be 'a' not '%s'\n", err.Error())
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}

	to, err = lexxx.GetNextToken()

	if to.Type != matchers.OPERATOR {
		t.Errorf("Token type should be OPERATOR not %s\n", to.Type)
	}
	if to.Value != "-" {
		t.Errorf("Token value should be '-' not '%s'\n", err.Error())
	}
	if err != nil {
		t.Errorf("This isn't an error situtation %s\n", err.Error())
	}
}

func BenchmarkLexx(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lexxx := Lexx{}
		lexxx.SetStringInput(`"Sed ut perspiciatis unde omnis iste natus error-sit @ voluptatem accusantium doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo inventore veritatis et quasi: architecto beatae vitae dicta sunt explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut odit aut fugit, sed quia consequuntur magni dolores eos qui ratione voluptatem sequi nesciunt$. Neque porro quisquam est, qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla pariatur? At vero eos et accusamus et iusto odio dignissimos ducimus qui blanditiis praesentium voluptatum deleniti atque corrupti & quos dolores et quas molestias excepturi sint occaecati cupiditate non provident, similique sunt in culpa qui officia deserunt mollitia animi, id est laborum et dolorum fuga. Et harum #@ quidem rerum facilis est et expedita distinctio. Nam libero tempore, cum soluta nobis est eligendi optio cumque nihil impedit quo minus id quod maxime placeat facere possimus, omnis voluptas assumenda est, omnis dolor repellendus. Temporibus autem quibusdam et aut officiis debitis aut rerum necessitatibus saepe eveniet ut et voluptates repudiandae sint et molestiae non recusandae. Itaque earum rerum hic tenetur a sapiente delectus, ut aut reiciendis voluptatibus maiores alias consequatur aut' perferendis doloribus asperiores repellat."`)
		lexxx.AddMatcher(matchers.StartWordMatcher)
		lexxx.AddMatcher(matchers.StartIntegerMatcher)
		lexxx.AddMatcher(matchers.StartFloatMatcher)
		lexxx.AddMatcher(matchers.StartWhitespaceMatcher)
		lexxx.AddMatcher(matchers.ConfigOperatorMatcher([]string{
			".",
			"!",
			"?",
			",",
			"'",
			"\"",
			"-",
			":",
			"[",
			"]",
			"(",
			")",
			"#",
			"*",
			"|",
			"_",
			";",
			"&",
			"/",
			"%",
			"@",
			"$",
		}))

		wordCount := 0
		integerCount := 0
		floatCount := 0
		symbolCount := 0
		totalCount := 0
		for {
			t, err := lexxx.GetNextToken()
			if err != nil {
				log.Fatal(err.Error())
			}
			if t == nil {
				fmt.Printf("Did not find tokens at %s\n", string(lexxx.State.CurrentText))
				break
			}
			totalCount++
			if t.Type == matchers.WORD {
				wordCount++
				continue
			} else if t.Type == matchers.INTEGER {
				integerCount++
				continue
			} else if t.Type == matchers.FLOAT {
				floatCount++
				continue
			} else if t.Type == matchers.WHITESPACE {
				continue
			} else if t.Type == matchers.OPERATOR {
				symbolCount++
				continue
			}
			if t.Value == "EOF" {
				break
			}
			log.Printf("Unkown tokens: %s\n", t.Value)
			break
		}
	}
}
