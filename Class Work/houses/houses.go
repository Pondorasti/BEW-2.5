package main

import (
	"fmt"

	"github.com/alexflint/go-arg"
)

type House struct {
	// numberOfRooms string
	City string
	// address       string
	Price string
}

func main() {
	fmt.Println("Hello world")

	var args House
	arg.MustParse(&args)
	fmt.Println(args.City, args.Price)
}
