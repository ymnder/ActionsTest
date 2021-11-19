package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println("---------------------------")

	i := flag.String("date", "", "date")
	s := flag.String("title", "", "title")
	b := flag.String("description", "", "author")

	flag.Parse()
	fmt.Println(*i, *s, *b)

	output, err := exec.Command("git", "switch", "master").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
	output, err = exec.Command("git", "switch", "-c", "tmp/a").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
	output, err = exec.Command("touch", "tmp.md").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
	output, err = exec.Command("git", "add", "tmp.md").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
	output, err = exec.Command("git", "commit", "-m", "add tmp.md").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}

	output, err = exec.Command("hub", "pull-request", "--draft", "-m", "test title").CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
}

func exitProcess(output []byte, err error) {
	fmt.Fprintf(os.Stderr, "failed executing %v\n", string(output))
	fmt.Fprintf(os.Stderr, "failed executing %v\n", err)
	os.Exit(1)
}
