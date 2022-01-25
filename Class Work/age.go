package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your date of birth: ")
	stringDob, _ := reader.ReadString('\n')
	fmt.Print("Enter your age: ")
	stringAge, _ := reader.ReadString('\n')

	dob, _ := strconv.Atoi(stringDob[:len(stringDob)-1])
	age, _ := strconv.Atoi(stringAge[:len(stringAge)-1])
	currentYear := dob + age
	fmt.Println("You are currently living in: ", currentYear)
}
