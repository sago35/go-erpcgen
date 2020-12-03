package main

import (
	"fmt"
	"strings"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
	Structs    map[string]*StructStatement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type OptionStatement struct {
	Token Token
	Name  *Identifier
	Value string
}

func (os *OptionStatement) statementNode() {}
func (os *OptionStatement) TokenLiteral() string {
	return os.Token.Literal
}

func (os *OptionStatement) String() string {
	return fmt.Sprintf("@%s(%q)", os.Name.Value, os.Value)
}

type ProgramStatement struct {
	Token Token
	Name  *Identifier
	Value Expression
}

func (ps *ProgramStatement) statementNode() {}
func (ps *ProgramStatement) TokenLiteral() string {
	return ps.Token.Literal
}

func (ps *ProgramStatement) String() string {
	return fmt.Sprintf("program %s", ps.Name.Value)
}

type EnumStatement struct {
	Token        Token
	Name         *Identifier
	Enumerations []Enumeration
}

func (es *EnumStatement) statementNode() {}
func (es *EnumStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *EnumStatement) String() string {
	ret := fmt.Sprintf("enum %s\n", es.Name.Value)
	ret += fmt.Sprintf("{\n")
	for _, enum := range es.Enumerations {
		ret += fmt.Sprintf("    %s\n", enum.String())
	}
	ret += fmt.Sprintf("}")
	return ret
}

type Enumeration struct {
	Name  string
	Value int64
	IsHex bool
}

func (e *Enumeration) String() string {
	if e.IsHex {
		return fmt.Sprintf("%s = 0x%X", e.Name, e.Value)
	} else {
		return fmt.Sprintf("%s = %d", e.Name, e.Value)
	}
}

type InterfaceStatement struct {
	Token      Token
	Name       *Identifier
	Interfaces []InterfaceDefinitionStatement
}

func (is *InterfaceStatement) statementNode() {}
func (is *InterfaceStatement) TokenLiteral() string {
	return is.Token.Literal
}

func (is *InterfaceStatement) String() string {
	ret := fmt.Sprintf("interface %s {\n", is.Name.Value)
	for _, ids := range is.Interfaces {
		ret += fmt.Sprintf("    %s\n", ids.String())
	}
	ret += fmt.Sprintf("}")
	return ret
}

type InterfaceDefinitionStatement struct {
	Token     Token
	Name      *Identifier
	Arguments []Argument
	Return    Argument
	Oneway    bool
}

func (ids *InterfaceDefinitionStatement) String() string {
	ret := fmt.Sprintf("%s(", ids.Name.Value)
	for i, arg := range ids.Arguments {
		if i == 0 {
			ret += fmt.Sprintf("%s", arg.String())
		} else {
			ret += fmt.Sprintf(", %s", arg.String())
		}
	}

	ret += fmt.Sprintf(")")
	if !ids.Oneway {
		ret += fmt.Sprintf(" -> ")
		for _, o := range ids.Return.Options {
			ret += fmt.Sprintf("@%s ", o)
		}
		ret += fmt.Sprintf("%s", strings.Join(ids.Return.Type, " "))
	}
	return ret
}

type Argument struct {
	Type      []string
	ArraySize string
	Name      string
	Options   []string
}

func (a *Argument) String() string {
	ret := fmt.Sprintf("%s", strings.Join(a.Type, " "))
	if a.ArraySize != "" {
		ret += fmt.Sprintf("[%s]", a.ArraySize)
	}
	ret += fmt.Sprintf(" %s", a.Name)

	for _, opt := range a.Options {
		ret += fmt.Sprintf(" @%s", opt)
	}
	return ret
}

func (os *InterfaceDefinitionStatement) statementNode() {}
func (os *InterfaceDefinitionStatement) TokenLiteral() string {
	return os.Token.Literal
}

type StructStatement struct {
	Token       Token
	Name        *Identifier
	Definitions []Definition
}

func (ss *StructStatement) statementNode() {}
func (ss *StructStatement) TokenLiteral() string {
	return ss.Token.Literal
}

func (ss *StructStatement) String() string {
	ret := fmt.Sprintf("struct %s\n", ss.Name.Token.Literal)
	ret += fmt.Sprintf("{\n")
	for _, def := range ss.Definitions {
		ret += fmt.Sprintf("    %s;\n", def.String())
	}
	ret += fmt.Sprintf("}")
	return ret
}

type Definition struct {
	Type      []string
	ArraySize string
	Name      string
}

func (d *Definition) String() string {
	if d.ArraySize == "" {
		return fmt.Sprintf("%s %s", strings.Join(d.Type, " "), d.Name)
	} else {
		return fmt.Sprintf("%s[%s] %s", strings.Join(d.Type, " "), d.ArraySize, d.Name)
	}
}
