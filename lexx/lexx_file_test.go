package lexx

import (
	"bufio"
	"fmt"
	"github.com/JeffThomas/lexx/matchers"
	"log"
	"os"
	"testing"
	"time"
)

func TestLexx_LoadFile(t *testing.T) {

	file, err := os.Open("../Varney-the-Vampire.txt")
	if err != nil {
		print(err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Printf("Closing file %s failed.", "../Varney-the-Vampire.txt")
			fmt.Println(err)
		}
	}(file)

	wordsMap := make(map[string]int)
	undefined := make(map[string]int)

	r := bufio.NewReader(file)

	now := time.Now()
	unixNano := now.UnixNano()
	umillisec := unixNano / 1000000

	lexxx := NewLexx([]matchers.LexxMatcherInitialize{
		matchers.StartWordMatcher,
		matchers.StartIntegerMatcher,
		matchers.StartFloatMatcher,
		matchers.StartWhitespaceMatcher,
		matchers.ConfigOperatorMatcher([]string{
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
		}),
	})
	lexxx.input = r

	wordCount := 0
	integerCount := 0
	floatCount := 0
	symbolCount := 0
	totalCount := 0
	lineCount := 0

	for {
		t, err := lexxx.GetNextToken()
		if err != nil {
			log.Fatal(err.Error())
		}
		if t == nil {
			fmt.Printf("Did not find tokens at %s\n", string(lexxx.State.CurrentText))
			break
		}
		//fmt.Printf("%s", t.Value)
		totalCount++
		lineCount = t.Line
		if t.Type == matchers.WORD {
			//fmt.Printf("%s\n", t.value)
			count, ok := wordsMap[t.Value]
			if ok {
				wordsMap[t.Value] = count + 1
			} else {
				wordsMap[t.Value] = 1
			}
			wordCount++
			continue
		} else if t.Type == matchers.INTEGER {
			//fmt.Printf("%s\n", t.value)
			integerCount++
			continue
		} else if t.Type == matchers.FLOAT {
			//fmt.Printf("%s\n", t.value)
			floatCount++
			continue
		} else if t.Type == matchers.WHITESPACE {
			// do nothing
			continue
		} else if t.Type == matchers.UNDEFINED {
			count, ok := undefined[t.Value]
			if ok {
				undefined[t.Value] = count + 1
			} else {
				undefined[t.Value] = 1
			}
			continue
		} else if t.Type == matchers.OPERATOR {
			symbolCount++
			continue
		} else if t.Type == matchers.SYSTEM && t.Value == "EOF" {
			log.Println("EndOf File.")
			break
		}
	}
	nowFinished := time.Now()
	unixNanoFinished := nowFinished.UnixNano()
	umillisecFinished := unixNanoFinished / 1000000
	fmt.Println("Elapsed milliseconds : ", umillisecFinished-umillisec)
	fmt.Println()
	fmt.Printf("Total tokens count %d\n", totalCount)
	fmt.Printf("Word count %d\n", wordCount)
	fmt.Printf("Unique word count %d\n", len(wordsMap))
	fmt.Printf("Integer count %d\n", integerCount)
	fmt.Printf("Float count %d\n", floatCount)
	fmt.Printf("Symbols count %d\n", symbolCount)
	fmt.Printf("Line Count %d\n", lineCount)
	fmt.Printf("Undefined Count %d\n", len(undefined))

	for k, v := range undefined {
		fmt.Printf("undefined '%s': %d\n", k, v)
	}

}
