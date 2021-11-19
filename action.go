package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

func main() {

	fmt.Println("---------------------------")
	err := errors.New("hello.")
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed executing command with error %v\n", err)
		os.Exit(1)
	}

	i := flag.String("date", "", "date")
	s := flag.String("title", "", "title")
	b := flag.String("description", "", "author")

	flag.Parse()
	fmt.Println(*i, *s, *b)
}
