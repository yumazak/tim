package internal

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// fixedOffsetSeconds returns the standard-time UTC offset in seconds for a timezone.
// Uses epoch (1970-01-01 00:00:00 UTC) to avoid DST.
func fixedOffsetSeconds(loc *time.Location) int {
	epoch := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	_, offset := epoch.In(loc).Zone()
	return offset
}

func convertHour(hour int, fromLoc, toLoc *time.Location) int {
	fromOff := fixedOffsetSeconds(fromLoc)
	toOff := fixedOffsetSeconds(toLoc)
	diffMinutes := (toOff - fromOff) / 60
	totalMinutes := hour*60 + diffMinutes
	h := totalMinutes / 60
	// Go's % can be negative, so use manual euclidean mod
	h = ((h % 24) + 24) % 24
	return h
}

// RunH handles the "h" subcommand.
func RunH(args []string) int {
	fs, from, to := NewZoneFlags("h")
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
		result, err := processHour(fs.Arg(0), fromLoc, toLoc)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return 1
		}
		fmt.Println(result)
		return 0
	}

	return ProcessStdin(func(line string) (string, error) {
		return processHour(line, fromLoc, toLoc)
	})
}

func processHour(input string, fromLoc, toLoc *time.Location) (string, error) {
	input = strings.TrimSpace(input)
	hour, err := strconv.Atoi(input)
	if err != nil || hour < 0 || hour > 23 {
		return "", fmt.Errorf("invalid hour: %s", input)
	}
	return strconv.Itoa(convertHour(hour, fromLoc, toLoc)), nil
}
