package main


import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check_number(reader bufio.Reader)(int){
	for {
		fmt.Print("Give an arabic number, I'll convert it for you: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input, please enter a valid number.")
			continue
		}

		if number == 0 {
			fmt.Println("The number must not be zero. Please try again.")
			continue
		}

		fmt.Println("You entered:", number)
		return number

	}
}

func treat(chiffre int){
	// 1 --> I
    // 10 --> X
     //7 --> VII
	if chiffre == 1 {
		println("I")
	}
	if (chiffre == 10){
		println("X")
	}
	if (chiffre == 5) {
		println("V")
	}

	
}
func main() {
	number := bufio.NewReader(os.Stdin)
	numInt := check_number(*number)
	treat(numInt)

}
