package internal

import (
	"flag"
	"fmt"
	"os"
)

// Run is the main entry point for the CLI.
func Run(args []string) int {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: tim <h|dt> [options] [value]")
		return 1
	}

	switch args[0] {
	case "h":
		return RunH(args[1:])
	case "dt":
		return RunDt(args[1:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", args[0])
		return 1
	}
}

// NewZoneFlags creates a FlagSet with --from/-f and --to/-t flags.
func NewZoneFlags(name string) (*flag.FlagSet, *string, *string) {
	fs := flag.NewFlagSet(name, flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	from := fs.String("from", "Asia/Tokyo", "source timezone (IANA name)")
	to := fs.String("to", "UTC", "target timezone (IANA name)")
	fs.StringVar(from, "f", "Asia/Tokyo", "source timezone (short)")
	fs.StringVar(to, "t", "UTC", "target timezone (short)")
	return fs, from, to
}
