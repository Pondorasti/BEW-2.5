package main

import (
	"fmt"
)

func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)
}
