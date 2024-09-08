package ui

import (
	"fmt"
	"os"
	"strings"

	fzf "github.com/junegunn/fzf/src"
)

func Display(lines []string) (int, error) {
	inputChan := make(chan string)
	go func() {
		for _, line := range lines {
			inputChan <- line
		}
		close(inputChan)
	}()

	outputChan := make(chan string)
	go func() {
		for s := range outputChan {
			out := strings.Split(s, " ")
			fmt.Fprintf(os.Stdout, "%s\n", strings.Trim(out[0], " "))
		}
	}()

	options, err := fzf.ParseOptions(
		true,
		[]string{"--height=40%"},
	)
	if err != nil {
		return fzf.ExitError, err
	}

	options.Input = inputChan
	options.Output = outputChan

	code, err := fzf.Run(options)

	return code, err
}
