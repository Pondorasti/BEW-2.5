package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter an array of weights: ")
	weights, _ := reader.ReadString('\n')

	var sum float64
	var count int
	for _, weight := range strings.Fields(weights) {
		weightFloat, _ := strconv.ParseFloat(weight, 64)
		sum += weightFloat
		count++
	}

	fmt.Println("Average weight: ", sum/float64(count))
}
