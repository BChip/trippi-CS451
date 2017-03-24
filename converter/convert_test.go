package converter

import (
	"testing"

	"github.com/bchip/trippi-CS451/lexer"
)

func TestConverter(t *testing.T) {

	input := `
	let numOne = 5; 
	let numTwo = 10; 
	let sum = numOne + numTwo; 
	print("The sum of num1 and num2 is: "+sum); 
	if(sum > 20){ 
		print("The sum is greater than 20"); 
	}
	`

	l := lexer.New(input)
	c := New(l)
	c.ConvertProgram()

}
