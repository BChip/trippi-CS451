package parser

/* NOTES:
<expr> → <term> {(+ | -) <term>}
<term> → <factor> {(* | /) <factor>}
<factor> → id | int_constant | ( <expr> )
*/

import (
	"fmt"

	"github.com/bchip/trippi-CS451/lexer"
	"github.com/bchip/trippi-CS451/token"
)

type Parser struct {
	l         *lexer.Lexer
	errors    []string
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	return p
}

func (p *Parser) ParseProgram() {
	for !p.peekTokenIs(token.EOF) {
		switch p.curToken.Type {
		case "LET":
			p.parseLetStatement()
		case "IF":
			p.parseIfExpression()
		case "WHILE":
			p.parseWhile()
		case "FOR":
			p.parseFor()
		}
		p.nextToken()

	}
}

func (p *Parser) parseLetStatement() {
	if !p.expectPeek(token.IDENT) {
		return
	}
	if !p.expectPeek(token.ASSIGN) {
		return
	}
	p.nextToken()
	if p.peekTokenIs(token.PLUS) || p.peekTokenIs(token.MINUS) {
		p.expr()
	} else if p.peekTokenIs(token.ASTERISK) || p.peekTokenIs(token.SLASH) {
		p.term()
	} else if p.peekTokenIs(token.SEMICOLON) {
		p.isLiteral()
	} else {
		msg := fmt.Sprintf("Incorrect let statement!")
		p.errors = append(p.errors, msg)
	}

}

func (p *Parser) expr() {
	for p.peekTokenIs(token.PLUS) || p.peekTokenIs(token.MINUS) {
		p.nextToken()
		p.term()
	}
}

func (p *Parser) term() {
	p.factor()
	for p.peekTokenIs(token.ASTERISK) || p.peekTokenIs(token.SLASH) {
		p.nextToken()
		p.factor()
	}
}

func (p *Parser) factor() {
	if p.peekTokenIs(token.IDENT) || p.peekTokenIs(token.INT) {
		p.nextToken()
	} else {
		if p.expectPeek(token.INT) {
			//p.nextToken()
			p.expr()
			if !p.expectPeek(token.INT) {
				return
			}
		} else {
			return
		}
	}
}

func (p *Parser) parseIfExpression() {
	if !p.expectPeek(token.LPAREN) {
		return
	}
	p.nextToken()
	p.boolExpr()
	if !p.expectPeek(token.RPAREN) {
		return
	}
	if !p.expectPeek(token.LBRACE) {
		return
	}
	p.ParseProgram()
	if !p.expectCur(token.RBRACE) {
		return
	}

}

func (p *Parser) parseWhile() {
	p.parseIfExpression()
}

func (p *Parser) parseFor() {
	if !p.expectPeek(token.LPAREN) {
		return
	}
	if !p.expectPeek(token.LET) {
		return
	}
	p.parseLetStatement()
	p.nextToken()
	p.nextToken()
	p.boolExpr()
	if !p.expectPeek(token.SEMICOLON) {
		return
	}
	if !p.expectPeek(token.IDENT) {
		return
	}

	if !p.peekTokenIs(token.INC) && !p.peekTokenIs(token.DEC) {
		msg := fmt.Sprintf("Missing ++ or --")
		p.errors = append(p.errors, msg)
		return
	}
	p.nextToken()
	if !p.expectPeek(token.RPAREN) {
		return
	}
	if !p.expectPeek(token.LBRACE) {
		return
	}
	p.ParseProgram()
	if !p.expectCur(token.RBRACE) {
		return
	}
}

func (p *Parser) boolExpr() {
	if p.peekTokenIs(token.LT) || p.peekTokenIs(token.GT) || p.peekTokenIs(token.EQ) || p.peekTokenIs(token.NOT_EQ) {
		if p.curTokenIs(token.IDENT) || p.curTokenIs(token.INT) {
			p.nextToken()
			if p.peekTokenIs(token.IDENT) || p.peekTokenIs(token.INT) {
				p.nextToken()
				return
			}
		}
	} else if p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
		return
	} else {
		msg := fmt.Sprintf("Incorrect if statement")
		p.errors = append(p.errors, msg)
	}
}

/* HELPERS */
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s but got %s", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) expectCur(t token.TokenType) bool {
	if p.curTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) isLiteral() {
	if p.curTokenIs(token.INT) || p.curTokenIs(token.IDENT) || p.curTokenIs(token.STRING) || p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
		return
	} else {
		msg := fmt.Sprintf("Expecting literal!")
		p.errors = append(p.errors, msg)
	}

}
