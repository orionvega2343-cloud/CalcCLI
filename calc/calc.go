package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage: calc <number> [<number> <number> ...]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

var (
	op = flag.String("op", "add", "operator [add/sub/mul/div]")
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("calc: ")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		usage()
	}

	a := args[0]
	b := args[1]

	parseA, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}

	parseB, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}

	switch *op {
	case "add":
		fmt.Printf("%d + %d = %d\n", parseA, parseB, parseA+parseB)

	case "sub":
		fmt.Printf("%d - %d = %d\n", parseA, parseB, parseA-parseB)

	case "mul":
		fmt.Printf("%d  * %d = %d\n", parseA, parseB, parseA*parseB)

	case "div":
		if parseB == 0 {
			log.Fatal("division by zero")
		}
		fmt.Printf("%d/%d = %d\n", parseA, parseB, parseA/parseB)

	default:
		usage()
	}

}
