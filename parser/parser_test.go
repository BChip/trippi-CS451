package parser

import (
	"fmt"
	"testing"

	"github.com/bchip/trippi-CS451/lexer"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"let x = 5 + 5;"},
		{"let y = true;"},
		{"if(4 != 5)"},
	}

	for _, test := range tests {

		l := lexer.New(test.input)
		p := New(l)
		p.ParseProgram()
		fmt.Println(p.errors)
	}
}
