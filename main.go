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
	flag.StringVar(&filePath, "I", "./erpc_idl/rpc_system.erpc", "Add search path for imports")
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

	if true {
		err = generateGoCode(program)
		if err != nil {
			return err
		}

	} else {
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
				fmt.Printf("enum _%s_ids\n", stmt.Name.Value)
				fmt.Printf("{\n")
				fmt.Printf("    k%s_service_id = %d,\n", stmt.Name.Value, sid)
				for j, x := range stmt.Interfaces {
					fmt.Printf("    k%s_%s_id = %d,\n", stmt.Name.Value, x.Name.Value, j+1)
				}
				fmt.Printf("};\n")
				sid++
			}
		}

		for _, stmt := range program.Statements {
			switch stmt := stmt.(type) {
			case *InterfaceStatement:
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

type GoArgument struct {
	In       bool
	Out      bool
	Name     string
	Typ      []string
	Nullable bool
}

func (a GoArgument) Type() string {
	return strings.Join(a.Typ, " ")
}

func (a GoArgument) Size() int {
	ret := 0
	switch a.Type() {
	case "uint8", "int8", "bool":
		ret = 1
	case "uint16", "int16":
		ret = 2
	case "uint32", "int32":
		ret = 4
	}
	if strings.HasPrefix(a.Type(), "RPC_T") {
		ret = 4
	}
	return ret
}

func (a GoArgument) String() string {
	str := []string{a.Name, ":"}
	if a.In && a.Out {
		str = append(str, "inout")
	} else if a.Out {
		str = append(str, "out")
	} else {
		str = append(str, "in")
	}

	str = append(str, a.Type())

	if a.Nullable {
		str = append(str, "nullable")
	}
	return strings.Join(str, " ")
}

func (a GoArgument) ArgString() string {
	if a.Out {
		switch a.Type() {
		case "uint8", "int8", "bool", "uint16", "int16", "uint32", "int32":
			return fmt.Sprintf("%s *%s", a.Name, a.Type())
		}
	}
	return fmt.Sprintf("%s %s", a.Name, a.Type())
}

func (a *GoArgument) SetNullable(opt []string) {
	f := a.Nullable
	for _, o := range opt {
		if o == "nullable" {
			f = true
		}
	}
	a.Nullable = f
}

func generateGoCode(p *Program) error {
	sid := 1

	fmt.Printf("// Automatically generated file. DO NOT EDIT.\n")
	fmt.Printf("// Generated by github.com/sago35/go-erpcgen.\n")
	fmt.Printf("\n")
	fmt.Printf("package rtl8720dn\n")
	fmt.Printf("\n")
	fmt.Printf("import (\n")
	fmt.Printf("	\"encoding/binary\"\n")
	fmt.Printf("	\"fmt\"\n")
	fmt.Printf(")\n")
	fmt.Printf("\n")

	for _, stmt := range p.Statements {
		switch stmt := stmt.(type) {
		case *InterfaceStatement:
			//if stmt.Name.Value != "rpc_wifi_tcpip" {
			//	sid++
			//	continue
			//}
			for j, x := range stmt.Interfaces {
				//if x.Name.Value != "rpc_tcpip_adapter_init" {
				//	continue
				//}
				typ := strings.Join(x.Return.Type, " ")
				if typ == "binary" {
					typ = "[]byte"
				}

				arguments := []GoArgument{}
				for _, arg := range x.Arguments {
					a := GoArgument{
						In: true,
					}
					for _, t := range arg.Type {
						if t == "in" || t == "out" || t == "inout" {
							if t == "in" {
								a.In = true
								a.Out = false
							} else if t == "inout" {
								a.In = true
								a.Out = true
							} else { // out
								a.In = false
								a.Out = true
							}
							continue
						}
						if t == "binary" {
							t = "[]byte"
							//a.Nullable = true
						}
						a.Typ = append(a.Typ, t)
					}
					a.Name = arg.Name
					if a.Name == "func" {
						a.Name = "fn"
					} else if a.Name == "len" {
						a.Name = "length"
					}
					a.SetNullable(arg.Options)
					arguments = append(arguments, a)
				}

				argStr := []string{}
				for _, arg := range arguments {
					argStr = append(argStr, arg.ArgString())
				}

				funcName := x.Name.Value
				if typ == "" || typ == "void" {
					fmt.Printf("func (r *rtl8720dn) %s(%s) {\n", funcName, strings.Join(argStr, ", "))
				} else {
					fmt.Printf("func (r *rtl8720dn) %s(%s) %s {\n", funcName, strings.Join(argStr, ", "), typ)
				}

				fmt.Printf("	if r.debug {\n")
				fmt.Printf("		fmt.Printf(\"%s()\\r\\n\")\n", x.Name.Value)
				fmt.Printf("	}\n")
				fmt.Printf("	msg := startWriteMessage(0x00, 0x%02X, 0x%02X, uint32(r.seq))\n", sid, j+1)
				fmt.Printf("\n")

				if len(arguments) > 0 {
					for _, a := range arguments {
						if !a.In {
							continue
						}
						name := a.Name
						if a.Out {
							switch a.Type() {
							case "uint8", "int8", "bool", "uint16", "int16", "uint32", "int32", "[]byte":
								name = "*" + name
							}
						}

						fmt.Printf("	// %s\n", a.String())
						if a.Size() > 0 {
							if a.Type() == "bool" {
								fmt.Printf("	if %s {\n", name)
								fmt.Printf("		msg = append(msg, 1)\n")
								fmt.Printf("	} else {\n")
								fmt.Printf("		msg = append(msg, 0)\n")
								fmt.Printf("	}\n")
							} else {
								for i := 0; i < a.Size(); i++ {
									fmt.Printf("	msg = append(msg, byte(%s>>%d))\n", name, i*8)
								}
							}
						} else if a.Type() == "string" || a.Type() == "[]byte" {
							if a.Nullable {
								fmt.Printf("	if len(%s) == 0 {\n", name)
								fmt.Printf("		msg = append(msg, 1)\n")
								fmt.Printf("	} else {\n")
								fmt.Printf("		msg = append(msg, 0)\n")
								fmt.Printf("		msg = append(msg, byte(len(%s)), byte(len(%s)>>8), byte(len(%s)>>16), byte(len(%s)>>24))\n", name, name, name, name)
								fmt.Printf("		msg = append(msg, []byte(%s)...)\n", name)
								fmt.Printf("	}\n")
							} else {
								fmt.Printf("	msg = append(msg, byte(len(%s)), byte(len(%s)>>8), byte(len(%s)>>16), byte(len(%s)>>24))\n", name, name, name, name)
								fmt.Printf("	msg = append(msg, []byte(%s)...)\n", name)
							}
						} else {
							fmt.Printf("	msg = append(msg, %s...)\n", name)
						}
					}
				}
				fmt.Printf("\n")

				fmt.Printf("	r.performRequest(msg)\n")
				fmt.Printf("\n")
				fmt.Printf("	r.read()\n")

				useWindex := false
				for _, a := range arguments {
					if a.Out {
						useWindex = true
					}
				}
				if typ != "" && typ != "void" {
					useWindex = true
				}
				if useWindex {
					fmt.Printf("	widx := 8\n")
				}

				if len(arguments) > 0 {
					for _, a := range arguments {
						if !a.Out {
							continue
						}
						name := a.Name
						switch a.Type() {
						case "uint8", "int8", "bool", "uint16", "int16", "uint32", "int32":
							name = "*" + name
						}

						fmt.Printf("	// %s\n", a.String())
						switch a.Type() {
						case "bool":
							fmt.Printf("	%s = payload[widx] != 0\n", name)
							fmt.Printf("	widx += 1\n")
						case "uint8":
							fmt.Printf("	%s = payload[widx]\n", name)
							fmt.Printf("	widx += 1\n")
						case "uint16":
							fmt.Printf("	%s = binary.LittleEndian.Uint16(payload[widx:])\n", name)
							fmt.Printf("	widx += 2\n")
						case "uint32":
							fmt.Printf("	%s = binary.LittleEndian.Uint32(payload[widx:])\n", name)
							fmt.Printf("	widx += 4\n")
						case "int8":
							fmt.Printf("	%s = int8(payload[widx])\n", name)
							fmt.Printf("	widx += 1\n")
						case "int16":
							fmt.Printf("	%s = int16(binary.LittleEndian.Uint16(payload[widx:]))\n", name)
							fmt.Printf("	widx += 2\n")
						case "int32":
							fmt.Printf("	%s = int32(binary.LittleEndian.Uint32(payload[widx:]))\n", name)
							fmt.Printf("	widx += 4\n")
						case "string", "[]byte":
							if a.Nullable {
								fmt.Printf("	%s_length := binary.LittleEndian.Uint32(payload[widx:])\n", a.Name)
								fmt.Printf("	widx += 1\n")
								fmt.Printf("	if %s_length == 1 {\n", a.Name)
								fmt.Printf("		%s_length = binary.LittleEndian.Uint32(payload[widx:])\n", a.Name)
								fmt.Printf("		widx += 4\n")
								fmt.Printf("	}\n")
							} else {
								fmt.Printf("	%s_length := binary.LittleEndian.Uint32(payload[widx:])\n", a.Name)
								fmt.Printf("	widx += 4\n")
							}

							fmt.Printf("	if %s_length > 0 {\n", a.Name)
							if a.Type() == "string" {
								fmt.Printf("		%s = string(payload[widx:widx+int(%s_length)])\n", name, a.Name)
							} else {
								fmt.Printf("		copy(%s, payload[widx:widx+int(%s_length)])\n", name, a.Name)
							}
							fmt.Printf("		widx += int(%s_length)\n", a.Name)
							fmt.Printf("	}\n")
						default:
							if a.Size() > 0 {
								fmt.Println("// not impl (a.Size() > 0)")
							} else {
								fmt.Println("// not impl")
							}
						}
					}
					fmt.Printf("\n")
				}

				if typ == "" || typ == "void" {
				} else if typ == "[]byte" {
					fmt.Printf("	var result []byte\n")
					fmt.Printf("	%s_length := binary.LittleEndian.Uint32(payload[widx:])\n", "result")
					fmt.Printf("	widx += 4\n")
					fmt.Printf("	copy(result, payload[widx:%s_length])\n", "result")
				} else if typ == "bool" {
					fmt.Printf("	var result %s\n", typ)
					fmt.Printf("	result = binary.LittleEndian.Uint32(payload[widx:]) == 1\n")
				} else if typ == "string" {
					fmt.Printf("	var result %s\n", typ)
					fmt.Printf("	%s_length := binary.LittleEndian.Uint32(payload[widx:])\n", "result")
					fmt.Printf("	widx += 4\n")
					fmt.Printf("	result = %s(payload[widx:widx+int(%s_length)])\n", typ, "result")
				} else if typ == "int32" || typ == "int16" || typ == "int8" {
					fmt.Printf("	var result %s\n", typ)
					fmt.Printf("	x := binary.LittleEndian.Uint32(payload[widx:])\n")
					fmt.Printf("	result = %s(x)\n", typ)
				} else {
					fmt.Printf("	var result %s\n", typ)
					fmt.Printf("	result = %s(binary.LittleEndian.Uint32(payload[widx:]))\n", typ)
				}

				fmt.Printf("\n")
				fmt.Printf("	r.seq++\n")
				if typ == "" || typ == "void" {
					fmt.Printf("	return\n")
				} else {
					fmt.Printf("	return result\n")
				}
				fmt.Printf("}\n")
				fmt.Printf("\n")
			}
			sid++
		}
	}
	return nil
}
