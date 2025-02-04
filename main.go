package main

import (
	"bufio"
	"os"

	"github.com/Hindmalti/kata-roman-numerals/treat"
)

func main() {
	number := bufio.NewReader(os.Stdin)
	numInt := treat.Check_number(*number)
	treat.Treat(numInt)

}
