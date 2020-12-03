package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {
	l         *Lexer
	curToken  Token
	peekToken Token

	errors []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParserProgram() *Program {
	program := &Program{}
	program.Statements = []Statement{}

	for p.curToken.Type != TOK_EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() Statement {
	fmt.Printf("*** %#v\n", p.curToken)
	switch p.curToken.Type {
	case TOK_AT:
		return p.parseOptionStatement()
	case TOK_PROGRAM:
		return p.parseProgramStatement()
	case TOK_ENUM:
		return p.parseEnumStatement()
	case TOK_INTERFACE:
		return p.parseInterfaceStatement()
	case TOK_STRUCT:
		return p.parseStructStatement()
	default:
		return nil
	}
}

func (p *Parser) parseOptionStatement() *OptionStatement {
	stmt := &OptionStatement{Token: p.curToken}
	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(TOK_LPAREN) {
		return nil
	}

	if !p.peekTokenIs(TOK_STRING) {
		return nil
	}

	p.nextToken()
	stmt.Value = p.curToken.Literal
	//stmt.Value = p.parseExpression(LOWEST)

	if !p.peekTokenIs(TOK_RPAREN) {
		return nil
	}

	return stmt
}

func (p *Parser) parseProgramStatement() *ProgramStatement {
	stmt := &ProgramStatement{Token: p.curToken}
	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	return stmt
}

func (p *Parser) parseEnumStatement() *EnumStatement {
	stmt := &EnumStatement{Token: p.curToken}
	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(TOK_LBRACE) {
		return nil
	}
	p.nextToken()

	for {
		if p.curTokenIs(TOK_RBRACE) {
			break
		}
		enums := p.parseEnumerations()
		stmt.Enumerations = enums
	}

	if !p.expectPeek(TOK_SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseEnumerations() []Enumeration {
	ret := []Enumeration{}
	val := int64(0)

	for !p.curTokenIs(TOK_RBRACE) {
		if !p.curTokenIs(TOK_IDENT) {
			return nil
		}
		enum := Enumeration{
			Name: p.curToken.Literal,
		}

		if p.peekTokenIs(TOK_ASSIGN) {
			p.nextToken()

			if !p.expectPeek(TOK_NUM) {
				return nil
			}

			cur := p.curToken.Literal
			if strings.HasPrefix(strings.ToLower(cur), "0x") {
				enum.IsHex = true
			}
			x, err := strconv.ParseInt(cur, 0, 64)
			if err != nil {
				return nil
			}
			val = x
		}
		enum.Value = val
		val++

		ret = append(ret, enum)
		p.nextToken()

		for p.curTokenIs(TOK_COMMA) || p.curTokenIs(TOK_COMMENT) {
			p.nextToken()
		}

		fmt.Printf("-- %s\n", enum.String())
		fmt.Printf("   %s\n", p.curToken.Literal)
	}

	return ret
}

func (p *Parser) parseInterfaceStatement() *InterfaceStatement {
	stmt := &InterfaceStatement{Token: p.curToken}

	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(TOK_LBRACE) {
		return nil
	}

	for {
		if p.peekTokenIs(TOK_RBRACE) {
			break
		}
		ids := p.parseInterfaceDefinitionStatement()
		stmt.Interfaces = append(stmt.Interfaces, *ids)
	}

	if !p.expectPeek(TOK_RBRACE) {
		return nil
	}

	//if !p.peekTokenIs(TOK_TYPE) {
	//	return nil
	//}

	return stmt
}

func (p *Parser) parseInterfaceDefinitionStatement() *InterfaceDefinitionStatement {
	for p.peekTokenIs(TOK_COMMENT) {
		p.nextToken()
	}

	stmt := &InterfaceDefinitionStatement{Token: p.curToken}

	if p.peekTokenIs(TOK_ONEWAY) {
		stmt.Oneway = true
		p.nextToken()
	}

	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(TOK_LPAREN) {
		return nil
	}
	p.nextToken()

	for !p.curTokenIs(TOK_RPAREN) {
		args := p.parseInterfaceDefinitionArguments()
		stmt.Arguments = args
	}

	if !stmt.Oneway {

		if !p.expectPeek(TOK_ARROW) {
			return nil
		}

		p.nextToken()
		for p.curTokenIs(TOK_AT) {
			p.expectPeek(TOK_IDENT)
			opt := p.curToken.Literal
			p.nextToken()

			if p.curTokenIs(TOK_LPAREN) {
				for !p.curTokenIs(TOK_RPAREN) {
					opt += p.curToken.Literal
					p.nextToken()
				}
				opt += p.curToken.Literal
				p.nextToken()
			}

			stmt.Return.Options = append(stmt.Return.Options, opt)
		}

		if !(p.curTokenIs(TOK_TYPE) || p.curTokenIs(TOK_IDENT)) {
			return nil
		}

		stmt.Return.Type = append(stmt.Return.Type, p.curToken.Literal)
	}

	if p.peekTokenIs(TOK_SEMICOLON) {
		p.nextToken()
	}

	fmt.Printf("-- %s\n", stmt.String())
	return stmt
}

func (p *Parser) parseInterfaceDefinitionArguments() []Argument {
	ret := []Argument{}
	for !p.curTokenIs(TOK_RPAREN) {
		arg := Argument{}
		for p.curTokenIs(TOK_TYPE) || (p.curTokenIs(TOK_IDENT) && (p.peekTokenIs(TOK_IDENT) || p.peekTokenIs(TOK_LBRACKET))) {
			arg.Type = append(arg.Type, p.curToken.Literal)
			p.nextToken()
		}

		if p.curTokenIs(TOK_LBRACKET) {
			if !p.expectPeek(TOK_NUM) {
				return nil
			}
			arg.ArraySize = p.curToken.Literal

			if !p.expectPeek(TOK_RBRACKET) {
				return nil
			}
			p.nextToken()
		}

		if !p.curTokenIs(TOK_IDENT) {
			return nil
		}
		arg.Name = p.curToken.Literal

		p.nextToken()

		for p.curTokenIs(TOK_AT) {
			p.expectPeek(TOK_IDENT)
			opt := p.curToken.Literal
			p.nextToken()

			if p.curTokenIs(TOK_LPAREN) {
				for !p.curTokenIs(TOK_RPAREN) {
					opt += p.curToken.Literal
					p.nextToken()
				}
				opt += p.curToken.Literal
				p.nextToken()
			}

			arg.Options = append(arg.Options, opt)
		}

		ret = append(ret, arg)
		if p.curTokenIs(TOK_COMMA) {
			p.nextToken()
		}

		if p.curTokenIs(TOK_RBRACE) {
			p.nextToken()
		}
	}
	return ret
}

func (p *Parser) parseStructStatement() *StructStatement {
	stmt := &StructStatement{Token: p.curToken}

	if !p.expectPeek(TOK_IDENT) {
		return nil
	}

	stmt.Name = &Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(TOK_LBRACE) {
		return nil
	}
	p.nextToken()

	for {
		if p.curTokenIs(TOK_RBRACE) {
			break
		}
		stmt.Definitions = p.parseDefinitions()
	}

	if !p.expectPeek(TOK_SEMICOLON) {
		return nil
	}

	return stmt
}

func (p *Parser) parseDefinitions() []Definition {
	ret := []Definition{}

	for !p.curTokenIs(TOK_RBRACE) {
		def := Definition{}

		for p.curTokenIs(TOK_TYPE) || (p.curTokenIs(TOK_IDENT) && !p.peekTokenIs(TOK_SEMICOLON)) {
			def.Type = append(def.Type, p.curToken.Literal)
			p.nextToken()
		}

		if p.curTokenIs(TOK_LBRACKET) {
			if !p.expectPeek(TOK_NUM) {
				return nil
			}
			def.ArraySize = p.curToken.Literal

			if !p.expectPeek(TOK_RBRACKET) {
				return nil
			}
			p.nextToken()
		}

		if !p.curTokenIs(TOK_IDENT) {
			return nil
		}
		def.Name = p.curToken.Literal
		p.nextToken()

		if !p.curTokenIs(TOK_SEMICOLON) {
			return nil
		}
		p.nextToken()

		for p.curTokenIs(TOK_COMMENT) {
			p.nextToken()
		}

		ret = append(ret, def)
	}
	return ret
}

func (p *Parser) curTokenIs(t TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}
