package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/alexflint/go-arg"
)

type House struct {
	NumberOfRooms []int    `arg:"-n,separate"`
	City          []string `arg:"-c,separate"`
	Address       []string `arg:"-a,separate"`
	Price         []int    `arg:"-p,separate"`
}

// Usage Example
// go run houses.go -c "San Francisco" -a "555 Post" -p 123456 -n 8 -c "Lakeover" -a "Structural Lane Nr. 38th" -p 90000 -n 4

func main() {
	var args House
	arg.MustParse(&args)

	w := tabwriter.NewWriter(os.Stdout, 2, 2, 2, ' ', 0)
	fmt.Fprintln(w, "Address\tCity\tRooms\tPrice")
	for i := 0; i < len(args.Address); i++ {
		fmt.Fprintf(w, "%s\t%s\t%d\t$%d\n", args.Address[i], args.City[i], args.NumberOfRooms[i], args.Price[i])
	}
	w.Flush()
}
