package internal

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var datetimeFormats = []string{
	"2006-01-02T15:04:05",
	"2006-01-02 15:04:05",
	"2006-01-02T15:04",
	"2006-01-02 15:04",
}

// RunDt handles the "dt" subcommand.
func RunDt(args []string) int {
	fs, from, to := NewZoneFlags("dt")
	if err := fs.Parse(args); err != nil {
		return 1
	}

	fromLoc, err := time.LoadLocation(*from)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid timezone: %s\n", *from)
		return 1
	}
	toLoc, err := time.LoadLocation(*to)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid timezone: %s\n", *to)
		return 1
	}

	if fs.NArg() > 0 {
		result, err := processDatetime(fs.Arg(0), fromLoc, toLoc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(result)
		return 0
	}

	return ProcessStdin(func(line string) (string, error) {
		return processDatetime(line, fromLoc, toLoc)
	})
}

func processDatetime(input string, fromLoc, toLoc *time.Location) (string, error) {
	input = strings.TrimSpace(input)

	// Try RFC3339 first (has offset info)
	if t, err := time.Parse(time.RFC3339, input); err == nil {
		converted := t.In(toLoc)
		return formatDatetime(converted), nil
	}

	// Try naive formats (interpret in fromLoc)
	for _, layout := range datetimeFormats {
		if t, err := time.ParseInLocation(layout, input, fromLoc); err == nil {
			converted := t.In(toLoc)
			return formatDatetime(converted), nil
		}
	}

	return "", fmt.Errorf("invalid datetime: %s", input)
}

func formatDatetime(t time.Time) string {
	return t.Format(time.RFC3339)
}
