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
		{"const x = 5 + 5; let y = true;"},
		{"let y = x;"},
		{"if( true ){let x = 5 + 5;} else {let x = input;}"},
		{"for(let x = 0; x < 5; x++){ let f = 3; }"},
		{"print(\"whatup\");"},
		{"concat(x,y); length(x);"},
		{"let x[] = {4,4,4,5,5,5,5,5,5,5,5,5,5,5,5,5,6,6,6,5,5,54,4,5};"},
		{"let numOne = 5.54;"},
		{"let numTwo = 10;"},
		{"let sum = numOne + numTwo;"},
		{"print(\"The sum of num1 and num2 is: \"+sum);"},
		{"if(sum > 20){ print(\"The sum is greater than 20\"); }"},
	}

	for _, test := range tests {

		l := lexer.New(test.input)
		p := New(l)
		p.ParseProgram()
		fmt.Println(test)
		fmt.Println(p.errors)
	}
}
