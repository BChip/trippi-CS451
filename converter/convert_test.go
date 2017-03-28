package converter

import (
	"fmt"
	"github.com/bchip/trippi-CS451/lexer"
	"github.com/bchip/trippi-CS451/parser"
	"testing"
)

func TestConverter(t *testing.T) {

	input := `
	let numOne = 5; 
	let numTwo = 10; 
	let sum = numOne + numTwo; 
	print("The sum of num1 and num2 is: " + sum); 
	if(sum > 20){ 
		print("The sum is greater than 20"); 
	}
	`

	l := lexer.New(input)
	p := parser.New(l)
	p.ParseProgram()
	if len(p.Errors()) > 0 {
		fmt.Println("ERRORS: ", p.Errors())
		t.Fail()
		return
	}
	l = lexer.New(input)
	c := New(l)
	c.ConvertProgram()

}
