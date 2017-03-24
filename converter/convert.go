package converter

import (
	"bufio"
	"fmt"
	"github.com/bchip/trippi-CS451/lexer"
	"github.com/bchip/trippi-CS451/token"
	"os"
	"os/exec"
)

var writer *bufio.Writer

type Converter struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Converter {
	c := &Converter{
		l: l,
	}
	fileHandle, _ := os.Create("output.go")
	writer = bufio.NewWriter(fileHandle)
	fmt.Fprintln(writer, "package main")
	fmt.Fprintln(writer, "import \"fmt\"")
	fmt.Fprintln(writer, "func main() {")
	return c
}

func (p *Converter) ConvertProgram() {

	for !p.peekTokenIs(token.EOF) {
		switch p.curToken.Type {
		case "IF":
			p.convertIfExpression()
		case "LET":
			p.convertLetStatement()
		case "PRINT":
			p.convertPrint()

		}
		p.nextToken()
	}
	fmt.Fprintln(writer, "}")
	writer.Flush()
	exec.Command("gofmt", "-w", "output.go").Run()
	out, _ := exec.Command("go", "run", "output.go").Output()
	fmt.Println(string(out))
}

func (p *Converter) convertLetStatement() {
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal+":=")
	p.nextToken()
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	if p.peekTokenIs(token.PLUS) {
		p.nextToken()
		fmt.Fprint(writer, p.curToken.Literal)
		p.nextToken()
		fmt.Fprint(writer, p.curToken.Literal)
	}
	fmt.Fprintln(writer)

}

func (p *Converter) convertIfExpression() {
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
	p.ConvertProgram()
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal)
}

func (p *Converter) convertPrint() {
	fmt.Fprint(writer, "fmt.Println(\"")
	p.nextToken()
	p.nextToken()
	fmt.Fprint(writer, p.curToken.Literal+"\"")
	p.nextToken()
	if p.curTokenIs(token.PLUS) {
		fmt.Fprint(writer, ",")
		p.nextToken()
		fmt.Fprint(writer, p.curToken.Literal)
		p.nextToken()
	}
	fmt.Fprint(writer, p.curToken.Literal)
	p.nextToken()
	fmt.Fprintln(writer)
}

func (p *Converter) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Converter) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Converter) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
