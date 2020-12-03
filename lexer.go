package main

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '@':
		tok = newToken(TOK_AT, l.ch)
	case ',':
		tok = newToken(TOK_COMMA, l.ch)
	case ';':
		tok = newToken(TOK_SEMICOLON, l.ch)
	case '=':
		tok = newToken(TOK_ASSIGN, l.ch)
	case '(':
		tok = newToken(TOK_LPAREN, l.ch)
	case ')':
		tok = newToken(TOK_RPAREN, l.ch)
	case '{':
		tok = newToken(TOK_LBRACE, l.ch)
	case '}':
		tok = newToken(TOK_RBRACE, l.ch)
	case '[':
		tok = newToken(TOK_LBRACKET, l.ch)
	case ']':
		tok = newToken(TOK_RBRACKET, l.ch)
	case '-':
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = Token{Type: TOK_ARROW, Literal: literal}
		}
	case '"':
		l.readChar()
		if isString(l.ch) {
			tok.Literal = l.readString()
			tok.Type = TOK_STRING
			l.readChar()
			return tok
		} else {
			tok.Literal = ""
			tok.Type = TOK_STRING
			l.readChar()
			return tok
		}
	case '/':
		if l.peekChar() == '/' {
			tok.Literal = l.readLine()
			tok.Type = TOK_COMMENT
		} else if l.peekChar() == '*' {
			tok.Literal = l.readComment()
			tok.Type = TOK_COMMENT
		}
	case 0:
		tok.Literal = ""
		tok.Type = TOK_EOF
	default:
		if l.ch == '0' && (l.peekChar() == 'x' || l.peekChar() == 'X') {
			tok.Type = TOK_NUM
			tok.Literal = l.readHexNumber()
			return tok
		} else if isDigit(l.ch) {
			tok.Type = TOK_NUM
			tok.Literal = l.readNumber()
			return tok
		} else if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = LookupIdent(tok.Literal)
			return tok
		}

	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readHexNumber() string {
	position := l.position
	l.readChar() // 0
	l.readChar() // x
	for isHexDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position
	for isString(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readLine() string {
	position := l.position
	for {
		if l.ch == '\n' || l.ch == '\r' {
			return l.input[position:l.position]
		}
		l.readChar()
	}
}

func (l *Lexer) readComment() string {
	position := l.position
	for {
		if l.ch == '*' && l.peekChar() == '/' {
			l.readChar()
			l.readChar()
			return l.input[position:l.position]
		}
		l.readChar()
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func newToken(tokenType TokenType, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || isDigit(ch)
}

func isString(ch byte) bool {
	return ch != '"'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isHexDigit(ch byte) bool {
	return '0' <= ch && ch <= '9' || 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F'
}
