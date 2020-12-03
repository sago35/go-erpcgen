package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
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
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	l := NewLexer(string(b))
	p := NewParser(l)

	program := p.ParserProgram()
	if program == nil {
		return fmt.Errorf("ParserProgram() returned nil")
	}

	sid := 1
	retMap := map[string]string{
		"":       "void",
		"string": "char *",
		"binary": "binary_t *",
		"int8":   "int8_t",
		"uint8":  "uint8_t",
		"int16":  "int16_t",
		"uint16": "uint16_t",
		"int32":  "int32_t",
		"uint32": "uint32_t",
	}
	argMap := map[string]string{
		"":       "void",
		"string": "char",
		"binary": "binary_t",
		"int8":   "int8_t",
		"uint8":  "uint8_t",
		"int16":  "int16_t",
		"uint16": "uint16_t",
		"int32":  "int32_t",
		"uint32": "uint32_t",
	}
	primitiveArgMap := map[string]string{
		"int8":   "int8_t",
		"uint8":  "uint8_t",
		"int16":  "int16_t",
		"uint16": "uint16_t",
		"int32":  "int32_t",
		"uint32": "uint32_t",
	}

	for _, stmt := range program.Statements {
		switch stmt := stmt.(type) {
		case *InterfaceStatement:
			if false {
				fmt.Printf("enum _%s_ids\n", stmt.Name.Value)
				fmt.Printf("{\n")
				fmt.Printf("    k%s_service_id = %d,\n", stmt.Name.Value, sid)
				for j, x := range stmt.Interfaces {
					fmt.Printf("    k%s_%s_id = %d,\n", stmt.Name.Value, x.Name.Value, j+1)
				}
				fmt.Printf("};\n")
			}
			sid++

			for _, x := range stmt.Interfaces {
				typ := strings.Join(x.Return.Type, " ")
				if rm, ok := retMap[typ]; ok {
					typ = rm
				}
				fmt.Printf("%s ", typ)

				if strings.Join(x.Return.Options, " ") == "nullable" {
					if !strings.HasSuffix(typ, " *") {
						fmt.Printf("* ")
					}
				}
				fmt.Printf("%s(", x.Name.Value)
				if len(x.Arguments) == 0 {
					fmt.Printf("void")
				} else {
					for i, arg := range x.Arguments {
						if i != 0 {
							fmt.Printf(", ")
						}
						typ := ""
						isPointer := false
						isConst := false
						isOut := false

						for _, t := range arg.Type {
							switch t {
							case "string":
								isPointer = true
								isConst = true
							case "binary":
								isPointer = true
								isConst = true
							}
							if _, ok := program.Structs[t]; ok {
								isPointer = true
								isConst = true
							}
						}

						for _, t := range arg.Type {
							switch t {
							case "in":
								isPointer = true
								isConst = true
							case "out":
								isPointer = true
								isOut = true
								isConst = false
							case "inout":
								isPointer = true
								isOut = true
								isConst = false
							default:
								if _, ok := primitiveArgMap[t]; ok {
									if !isOut {
										isPointer = false
									}
									isConst = false
								}

								if typ == "" {
									typ = t
								} else {
									typ += fmt.Sprintf(" %s", t)
								}
							}
						}

						for _, o := range arg.Options {
							switch o {
							case "nullable":
								//isConst = false
							}
						}

						sz := arg.ArraySize
						if sz != "" {
							if !isOut {
								isConst = true
							}
							isPointer = false
						}

						// struct or enum で処理が変わる
						if am, ok := argMap[typ]; ok {
							typ = am
						}

						// 基本型の場合は const * つけない
						if isPointer && isConst {
							fmt.Printf("const %s * %s", typ, arg.Name)
						} else if isConst {
							fmt.Printf("const %s %s", typ, arg.Name)
						} else if isPointer {
							fmt.Printf("%s * %s", typ, arg.Name)
						} else {
							fmt.Printf("%s %s", typ, arg.Name)
						}

						if sz != "" {
							fmt.Printf("[%s]", sz)
						}
					}
				}
				fmt.Printf(");\n")
			}
		default:
		}
		if false {
			fmt.Printf("--\n")
			fmt.Printf("%s\n", stmt)
		}
	}

	return nil
}
func run2(filePath string) error {
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
			fmt.Printf("enum _%s_ids\n", n.Literal)
			fmt.Printf("{\n")
			fmt.Printf("    k%s_service_id = %d,\n", n.Literal, interfaceNo)

			// skip {
			i++

			for tokens[i].Type != TOK_RBRACE {
				if interfaceNo == 1 {
					if tokens[i].Type == TOK_IDENT {
						fmt.Printf("    %#v\n", tokens[i])
					}
				}
				i++
			}
			fmt.Printf("}\n")
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
