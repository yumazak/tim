package internal

import (
	"bufio"
	"fmt"
	"os"
)

// ProcessStdin reads lines from stdin, processes each with the given function,
// and writes results to stdout. Errors are written to stderr and processing continues.
func ProcessStdin(process func(string) (string, error)) int {
	scanner := bufio.NewScanner(os.Stdin)
	hadError := false
	for scanner.Scan() {
		line := scanner.Text()
		result, err := process(line)
		if err != nil {
			hadError = true
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "io error: %v\n", err)
		hadError = true
	}
	if hadError {
		return 1
	}
	return 0
}
