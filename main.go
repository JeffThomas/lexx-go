// Package main is the CLI for Lexx
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/JeffThomas/lexx/lexx"
	matchersclass "github.com/JeffThomas/lexx/matchers"
	"log"
	"os"
)

func main() {
	filePtr := flag.String("f", "", "File to tokenize")
	undefinedOnlyPtr := flag.Bool("u", false, "Only tokenize UNDEFINED (handy for debugging)")
	whitespacePtr := flag.Bool("w", false, "Include WHITESPACE in output")
	dividerPtr := flag.String("d", " ", "Divider for fields, default ' '")
	jsonPtr := flag.Bool("j", false, "Output in JSON format")

	flag.Parse()

	var r *bufio.Reader = nil
	if *filePtr != "" {
		file, err := os.Open(*filePtr)
		if err != nil {
			print(err)
			return
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Printf("Failed to close file %s.\n", *filePtr)
			}
		}(file)

		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(os.Stdin)
	}

	lexxx := lexx.NewDefaultLexx()
	lexxx.SetInput(r)

	if *jsonPtr {
		fmt.Printf("[")
	}

	for {
		t, err := lexxx.GetNextToken()
		if err != nil {
			log.Fatal(err.Error())
		}
		if !*whitespacePtr && t.Type == matchersclass.WHITESPACE {
			continue
		}
		if *undefinedOnlyPtr {
			if t.Type != matchersclass.UNDEFINED && !(t.Type == matchersclass.SYSTEM && t.Value == "EOF") {
				continue
			}
		}

		if *jsonPtr {
			m, er := json.Marshal(t)
			if er != nil {
				fmt.Printf("error")
			}
			if t.Type == matchersclass.SYSTEM && t.Value == "EOF" {
				fmt.Printf("%s", string(m))
			} else {
				fmt.Printf("%s,", string(m))
			}
		} else {
			fmt.Printf("%s%s%d%s%d%s%s\n", t.Type, *dividerPtr, t.Line+1, *dividerPtr, t.Column, *dividerPtr, t.Value)
		}
		if t.Type == matchersclass.SYSTEM && t.Value == "EOF" {
			break
		}
		//if t.Type == matchers_class.WORD {
		//	fmt.Printf("WORD %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.INTEGER {
		//	fmt.Printf("INTEGER %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.FLOAT {
		//	fmt.Printf("FLOAT %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.WHITESPACE {
		//	// do nothing
		//	fmt.Printf("WHITESPACE %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.UNDEFINED {
		//	fmt.Printf("UNDEFINED %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.SYMBOLIC {
		//	fmt.Printf("SYMBOLIC %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.PUNCTUATION {
		//	fmt.Printf("PUNCTUATION %s\n", t.Value)
		//	continue
		//} else if t.Type == matchers_class.STRING {
		//	fmt.Printf("STRING \"%s\"\n", strings.ReplaceAll(t.Value, `"`, `\\"`))
		//	continue
		//} else if t.Type == matchers_class.SYSTEM && t.Value == "EOF" {
		//	fmt.Printf("SYSTEM %s\n", t.Value)
		//	break
		//}
		//break
	}
	if *jsonPtr {
		fmt.Printf("]")
	}
}
