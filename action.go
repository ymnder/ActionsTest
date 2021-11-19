package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

const (
	templateFile = "./sample.md"
	outputDir    = "./"
	branchPrefix = "tmp/"
)

func main() {
	dateArg := flag.String("date", "", "date")
	titleArg := flag.String("title", "", "title")
	authorArg := flag.String("author", "", "author")

	flag.Parse()
	fmt.Println(*dateArg, *titleArg, *authorArg)

	publishDate := parseDate(*dateArg)
	fileError := createFile(publishDate, *titleArg, *authorArg)
	if fileError != nil {
		exitProcessWithError(fileError)
		return
	}

	createBranch(publishDate)
}

func parseDate(inputDate string) string {
	dateRegex := regexp.MustCompile("^12/(0[1-9]|1[0-9]|2[0-5])$")
	publishDate := dateRegex.FindStringSubmatch(inputDate)
	if len(publishDate) == 0 {
		exitProcessWithMessage("\u2717 12/01のような形式で日付を入力してください")
	}
	return publishDate[1]
}

func createFile(publishDate string, title string, author string) error {
	publishDateForLabel := publishDate
	if strings.HasPrefix(publishDate, "0") {
		publishDateForLabel = publishDate[1:]
	}

	template, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return err
	}

	article := strings.Replace(string(template), "[date]", publishDate, -1)
	article = strings.Replace(article, "[date_label]", publishDateForLabel, -1)
	if len(title) > 0 {
		article = strings.Replace(article, "[title]", title, -1)
	}
	if len(author) > 0 {
		article = strings.Replace(article, "[author]", author, -1)
	}

	err = ioutil.WriteFile(outputDir+publishDate+".md", []byte(article), 0644)
	if err != nil {
		return err
	}
	fmt.Println("\u2714 ファイルを作成しました")
	fmt.Println(article)
	return nil
}

func runCommand(name string, args ...string) {
	fmt.Println(exec.Command(name, args...).String())
	output, err := exec.Command(name, args...).CombinedOutput()
	if err != nil {
		exitProcess(output, err)
	}
}

func createBranch(publishDate string) {
	targetBranch := branchPrefix + publishDate
	runCommand("git", "config", "user.name", "[bot] ${GITHUB_ACTOR}")
	runCommand("git", "config", "user.email", "${GITHUB_ACTOR}@users.noreply.github.com")
	runCommand("git", "switch", "-c", targetBranch)
	runCommand("git", "add", outputDir+publishDate+".md")
	runCommand("git", "commit", "-m", "Add a template")
	runCommand("git", "push", "origin", targetBranch)
	runCommand("gh", "pr", "create", "--title", publishDate+" Article", "--base", "${GITHUB_REF_NAME}")
}

func exitProcessWithMessage(message string) {
	fmt.Fprintf(os.Stderr, message)
	os.Exit(1)
}

func exitProcessWithError(err error) {
	fmt.Fprintf(os.Stderr, "failed executing %v\n", err)
	os.Exit(1)
}

func exitProcess(output []byte, err error) {
	fmt.Fprintf(os.Stderr, "failed executing %v\n", string(output))
	fmt.Fprintf(os.Stderr, "failed executing %v\n", err)
	os.Exit(1)
}
