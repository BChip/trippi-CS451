package parser

import (
	"fmt"

	"github.com/bchip/trippi-CS451/lexer"
	"github.com/bchip/trippi-CS451/token"
)


type Parser struct {
	l      *lexer.Lexer
	errors []string
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
	for p.curToken.Type != token.EOF {
		switch p.curToken.Type{
		case "LET":
			p.parseLetStatement()
		}
		p.nextToken()

	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseLetStatement() {
	p.expectPeek(token.IDENT)
	p.expectPeek(token.ASSIGN)
	p.nextToken()
    switch true{
    case p.curTokenIs(token.LBRACE):
	    p.handleEquation()
	    break;
    case p.curTokenIs(token.INT):
		break;
    case p.curTokenIs(token.TRUE):
		break
    case p.curTokenIs(token.FALSE):
	    break
    case p.curTokenIs(token.IDENT):
	    break
    default:
	    msg := fmt.Sprintf("expected next token to be a number, string, or int but got %s", p.peekToken.Type)
	    p.errors = append(p.errors, msg)
    }
	for !(p.curTokenIs(token.SEMICOLON)) {
		p.nextToken()
	}


}

func (p *Parser) handleEquation(){
	p.nextToken()
	for !(p.curTokenIs(token.RBRACE)){
		switch true{
		case p.curTokenIs(token.INT):
			if !(p.peekTokenIs(token.PLUS) && p.peekTokenIs(token.MINUS) && p.peekTokenIs(token.ASTERISK) && p.peekTokenIs(token.SLASH)){
				msg := fmt.Sprintf("expected next token to be a operator but got %s", p.peekToken.Type)
				p.errors = append(p.errors, msg)
				break;
			}
			break;
		default:
			msg := fmt.Sprintf("expected next token to be a number but got %s", p.curToken.Type)
			p.errors = append(p.errors, msg)
		}
		p.nextToken()
	}
	p.nextToken()
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

func (p *Parser) parseIfExpression()  {
	if !p.expectPeek(token.LPAREN) {
		return
	}

	p.nextToken()

	if !p.expectPeek(token.RPAREN) {
		return
	}
	if !p.expectPeek(token.LBRACE) {
		return
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()
		if !p.expectPeek(token.LBRACE) {
			return
		}
	}

}

