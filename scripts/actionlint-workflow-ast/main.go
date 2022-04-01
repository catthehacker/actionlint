package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/catthehacker/actionlint"
	"github.com/kr/pretty"
)

func run(args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	var src []byte
	var err error
	if len(args) <= 1 {
		src, err = ioutil.ReadAll(stdin)
	} else {
		if args[1] == "-h" || args[1] == "-help" || args[1] == "--help" {
			fmt.Fprintln(stdout, "Usage: go run ./scripts/actionlint-workflow-ast {workflow_file}")
			return 0
		}
		src, err = ioutil.ReadFile(args[1])
	}
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	w, errs := actionlint.Parse(src)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Fprintln(stderr, err)
		}
		return 1
	}
	pretty.Fprintf(stdout, "%# v\n", w)
	return 0
}

func main() {
	os.Exit(run(os.Args, os.Stdin, os.Stdout, os.Stderr))
}
