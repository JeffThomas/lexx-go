package lexx

import (
	"bufio"
	"fmt"
	matchers_class "github.com/JeffThomas/lexx/matchers"
	"log"
	"os"
	"testing"
	"time"
)

func TestLexx_LoadFile(t *testing.T) {
	now := time.Now()
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000

	file, err := os.Open("../../Varney-the-Vampire.txt")
	if err != nil {
		print(err)
		return
	}
	defer file.Close()

	wordsMap := make(map[string]int)

	r := bufio.NewReader(file)
	c, _, err := r.ReadRune()
	fmt.Printf("---> %d\n", c)
	lexxx := Lexx{}
	lexxx.Input = r
	lexxx.AddMatcher(matchers_class.StartWordMatcher)
	lexxx.AddMatcher(matchers_class.StartIntegerMatcher)
	lexxx.AddMatcher(matchers_class.StartFloatMatcher)
	lexxx.AddMatcher(matchers_class.StartWhitespaceMatcher)
	lexxx.AddMatcher(matchers_class.InitSymbolMatcher([]string{
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
		"\xef",
		"\xbb",
		"\xbf",
		"\xe2",
		"\x80",
		"\x9c",
		"\x9d",
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
			fmt.Printf("Did not find token at %s\n", lexxx.CurrentText)
			break
		}
		totalCount++
		if t.Type == matchers_class.WORD {
			//fmt.Printf("%s\n", t.value)
			count, ok := wordsMap[t.Value]
			if ok {
				wordsMap[t.Value] = count + 1
			} else {
				wordsMap[t.Value] = 1
			}
			wordCount++
			continue
		} else if t.Type == matchers_class.INTEGER {
			//fmt.Printf("%s\n", t.value)
			integerCount++
			continue
		} else if t.Type == matchers_class.FLOAT {
			//fmt.Printf("%s\n", t.value)
			floatCount++
			continue
		} else if t.Type == matchers_class.WHITESPACE {
			// do nothing
			continue
		} else if t.Type == matchers_class.SYMBOL {
			symbolCount++
			continue
		}
		log.Printf("Unkown token: %s\n", t.Value)
		break
	}

	fmt.Printf("Word count %d\n", wordCount)
	fmt.Printf("Unique word count %d\n", len(wordsMap))
	fmt.Printf("Integer count %d\n", integerCount)
	fmt.Printf("Float count %d\n", floatCount)
	fmt.Printf("Symbols count %d\n", symbolCount)
	fmt.Printf("Total tokens %d\n", totalCount)
	fmt.Printf("Line Count %d\n", lexxx.Line)

	//for k, v := range wordsMap {
	//	fmt.Printf("%s: %d\n", k, v)
	//}

	nowFinished := time.Now()
	unixNanoFinished := nowFinished.UnixNano()
	umillisecFinished := unixNanoFinished / 1000000
	fmt.Println("Elapsed miliseconds : ", umillisecFinished-umillisec)
}
