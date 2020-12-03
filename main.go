package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {
	filePath := ""
	flag.StringVar(&filePath, "I", ".", "Add search path for imports")
	flag.Parse()

	err := run(filePath)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	interfaceNo = 0
)

func run(filePath string) error {
	tokens, err := tokenize(filePath)
	if err != nil {
		return err
	}

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		switch token.Type {
		case TOK_INTERFACE:
			n := tokens[i+1]
			i++
			if n.Type != TOK_IDENT {
				return fmt.Errorf("interface: got %q, want %q", n.Type, TOK_STRING)
			}
			interfaceNo++
			fmt.Printf("k%s = %d\n", n.Literal, interfaceNo)

			// skip {
			i++

			for tokens[i].Type != TOK_RBRACE {
				if interfaceNo == 1 {
					if tokens[i].Type == TOK_IDENT {
						fmt.Printf("%#v\n", tokens[i])
					}
				}
				i++
			}
		}
	}
	return nil
}

func tokenize(filePath string) ([]Token, error) {
	ret := []Token{}

	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	l := NewLexer(string(b))

FOR_LOOP:
	for {
		token := l.NextToken()

		switch token.Type {
		case TOK_EOF:
			break FOR_LOOP
		case TOK_IMPORT:
			n := l.NextToken()
			if n.Type != TOK_STRING {
				return nil, fmt.Errorf("import: got %q, want %q", n.Type, TOK_STRING)
			}

			p := filepath.Join(filepath.Dir(filePath), n.Literal)
			tokens, err := tokenize(p)
			if err != nil {
				return nil, err
			}
			ret = append(ret, tokens...)
			continue
		default:
		}
		ret = append(ret, token)
		//fmt.Printf("%#v\n", token)
	}

	return ret, nil
}
