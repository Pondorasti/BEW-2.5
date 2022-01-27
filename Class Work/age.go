package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your month of birth (as a number): ")
	stringMonth, _ := reader.ReadString('\n')
	fmt.Print("Enter your day of birth: ")
	stringDay, _ := reader.ReadString('\n')
	fmt.Print("Enter your age: ")
	stringAge, _ := reader.ReadString('\n')

	month, _ := strconv.Atoi(stringMonth[:len(stringMonth)-1])
	day, _ := strconv.Atoi(stringDay[:len(stringDay)-1])
	age, _ := strconv.Atoi(stringAge[:len(stringAge)-1])

	now := time.Now()

	if month > int(now.Month()) || (month == int(now.Month()) && day >= now.Day()) {
		fmt.Println("hello")
		fmt.Println("You were born:", now.Year()-age-1)
	} else {
		fmt.Println("You were born:", now.Year()-age)
	}
}
