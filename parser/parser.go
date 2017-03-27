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
		case "}":
			return
		case "LET":
			p.parseLetStatement()
		case "IF":
			p.parseIfExpression()
		case "ELSE":
			msg := fmt.Sprintf("You cannot have an else statement if there is no if statement")
			p.errors = append(p.errors, msg)
			return
		case "WHILE":
			p.parseWhile()
		case "FOR":
			p.parseFor()
		case "CONST":
			p.parseLetStatement()
		case "PRINT":
			p.parsePrint()
		case "INPUT":
			p.parseLetStatement() //parseLetStatement
		case "CONCAT":
			p.parseConcat()
		case "LENGTH":
			p.parseLength()
		}
		p.nextToken()

	}
}

func (p *Parser) parseLetStatement() {
	if !p.expectPeek(token.IDENT) {
		return
	}
	if p.peekTokenIs(token.ARR) {
		p.nextToken()
		p.parseArr()
		return
	}
	if !p.expectPeek(token.ASSIGN) {
		return
	}
	if p.peekTokenIs(token.INPUT) {
		p.nextToken()
		if !p.expectPeek(token.SEMICOLON) {
			return
		}
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

func (p *Parser) parseArr() {
	if !p.expectPeek(token.ASSIGN) {
		return
	}
	if !p.expectPeek(token.LBRACE) {
		return
	}
	p.nextToken()
	if p.peekTokenIs(token.RBRACE) {
		p.isLiteral()
		p.nextToken()
		if !p.expectPeek(token.SEMICOLON) {
			return
		}
	} else if p.peekTokenIs(token.COMMA) {
		for !p.peekTokenIs(token.RBRACE) {
			p.isLiteral()
			if !p.expectPeek(token.COMMA) {
				return
			}
			p.nextToken()
		}
		p.nextToken()
		if !p.expectPeek(token.SEMICOLON) {
			return
		}
	} else {
		msg := fmt.Sprintf("Incorrect Array Initializer!")
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
	if p.peekTokenIs(token.IDENT) || p.peekTokenIs(token.INT) || p.peekTokenIs(token.FLOAT) {
		p.nextToken()
	} else {
		if p.expectPeek(token.INT) || p.expectPeek(token.FLOAT) {
			//p.nextToken()
			p.expr()
			if !p.expectPeek(token.INT) && !p.expectPeek(token.FLOAT) {
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
	if p.curTokenIs(token.ELSE) {
		p.parseElse()
	}

}

func (p *Parser) parseElse() {
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

func (p *Parser) parsePrint() {
	if !p.expectPeek(token.LPAREN) {
		return
	}
	p.nextToken()
	p.isLiteral()
	if p.peekTokenIs(token.PLUS) {
		for !p.peekTokenIs(token.RPAREN) {
			p.isLiteral()
			if !p.expectPeek(token.PLUS) {
				return
			}
			p.nextToken()
		}
		p.nextToken()
		if !p.expectPeek(token.SEMICOLON) {
			return
		}
	}
}

func (p *Parser) parseConcat() {
	if !p.expectPeek(token.LPAREN) {
		return
	}
	if !p.expectPeek(token.IDENT) && !p.expectPeek(token.STRING) {
		return
	}
	if !p.expectPeek(token.COMMA) {
		return
	}
	if !p.expectPeek(token.IDENT) && !p.expectPeek(token.STRING) {
		return
	}
	if !p.expectPeek(token.RPAREN) {
		return
	}
	if !p.expectPeek(token.SEMICOLON) {
		return
	}
}

func (p *Parser) parseLength() {
	if !p.expectPeek(token.LPAREN) {
		return
	}
	if !p.expectPeek(token.STRING) {
		return
	}
	if !p.expectPeek(token.RPAREN) {
		return
	}
	if !p.expectPeek(token.SEMICOLON) {
		return
	}
}

func (p *Parser) boolExpr() {
	if p.peekTokenIs(token.LT) || p.peekTokenIs(token.GT) || p.peekTokenIs(token.EQ) || p.peekTokenIs(token.NOT_EQ) {
		if p.curTokenIs(token.IDENT) || p.curTokenIs(token.INT) || p.curTokenIs(token.FLOAT) {
			p.nextToken()
			if p.peekTokenIs(token.IDENT) || p.peekTokenIs(token.INT) || p.peekTokenIs(token.FLOAT) {
				p.nextToken()
				return
			}
		}
	} else if p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
		return
	} else {
		msg := fmt.Sprintf("Incorrect statement")
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
	if p.curTokenIs(token.INT) || p.curTokenIs(token.FLOAT) || p.curTokenIs(token.IDENT) || p.curTokenIs(token.STRING) || p.curTokenIs(token.TRUE) || p.curTokenIs(token.FALSE) {
		return
	} else {
		msg := fmt.Sprintf("Expecting literal!")
		p.errors = append(p.errors, msg)
	}

}
