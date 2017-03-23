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
		{"const x = 5+ 5; let y = true;"},
		{"let y = x;"},
		{"if( true ){let x = 5 + 5;} else {let x = input;}"},
		{"for(let x = 0; x < 5; x++){ let f = 3; }"},
		{"print(\"whatup\");"},
		{"concat(x,y); length(x); else{}"},
	}

	for _, test := range tests {

		l := lexer.New(test.input)
		p := New(l)
		p.ParseProgram()
		fmt.Println(p.errors)
	}
}
