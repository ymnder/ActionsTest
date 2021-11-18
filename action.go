package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {

	fmt.Println("---------------------------")

	i := flag.String("date", "", "date")
	s := flag.String("title", "", "title")
	b := flag.String("description", "", "author")

	flag.Parse()
	fmt.Println(*i, *s, *b)
}
