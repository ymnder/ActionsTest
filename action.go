package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	dateArg := flag.String("date", "", "date")
	titleArg := flag.String("title", "", "title")
	authorArg := flag.String("author", "", "author")

	flag.Parse()
	fmt.Println(*dateArg, *titleArg, *authorArg)

	runCommand("git", "switch", "-c", "tmp/a")
	runCommand("touch", "tmp.md")
	runCommand("git", "add", "tmp.md")
	runCommand("git", "commit", "-m", "add tmp.md")
	runCommand("hub", "pull-request", "--draft", "-m", "test title")
}

func runCommand(name string, args ...string) {
	output, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
}

func exitProcess(output []byte, err error) {
	fmt.Fprintf(os.Stderr, "failed executing %v\n", string(output))
	fmt.Fprintf(os.Stderr, "failed executing %v\n", err)
	os.Exit(1)
}
