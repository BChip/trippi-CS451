package parser

import (
	"fmt"
	"testing"

	"github.com/bchip/trippi-CS451/lexer"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
	}{
		{"let x = 5 + 5;"},
		{"let y = true;"},
		{"let foobar = y;"},
	}

	for _, test := range tests {

		l := lexer.New(test.input)
		p := New(l)
		p.ParseProgram()
		fmt.Println(p.errors)
	}
}

