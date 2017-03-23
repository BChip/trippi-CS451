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
		{"let x = 5 + 5; let y = 4;"},
		{"let y = true;"},
		{"if(true){let x = 5 + 5;}"},
		{"for(let x = 0; x < 5; x++){ let f = 3; }"},
	}

	for _, test := range tests {

		l := lexer.New(test.input)
		p := New(l)
		p.ParseProgram()
		fmt.Println(p.errors)
	}
}
