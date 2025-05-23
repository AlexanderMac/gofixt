package main

import (
	"flag"
	"log"
	"os"
	"slices"

	"github.com/alexandermac/gofixt/internal"
)

const VERSION = "0.2.0"

func main() {
	log.SetFlags(0)

	flags := flag.NewFlagSet("gofixt", flag.ExitOnError)
	flags.Usage = usage
	dir := flags.String("dir", "", "Scanning directory")
	exts := flags.String("exts", "", "Extension list")
	printMode := flags.String("print", "important", "Print mode")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Unable to parse args: %v", err)
	}

	args := flags.Args()
	if len(args) < 1 {
		flags.Usage()
		os.Exit(1)
	}

	firstArg := args[0]
	switch firstArg {
	case "help":
		flags.Usage()
		os.Exit(0)
	case "version":
		log.Printf("v%s\n", VERSION)
		os.Exit(0)
	case "scan":
		validateFlags(*dir, *printMode)
		if err := internal.Scan(*dir, *exts, internal.PrintMode(*printMode)); err != nil {
			log.Fatal(err)
		}
	case "fix":
		validateFlags(*dir, *printMode)
		if err := internal.Fix(*dir, *exts, internal.PrintMode(*printMode)); err != nil {
			log.Fatal(err)
		}
	default:
		flags.Usage()
		os.Exit(0)
	}
}

func usage() {
	const usagePrefix = `Usage: gofixt [flags] command

Flags:
  --dir    Scanning directory (absolute or relative path)
  --exts   Comma separated list of file extensions, files with other extensions will be ignored, default: empty
  --print  Print mode: all,important,report,none, default: important

Commands:
  scan     Scan files in the provided directory recursively. Print report in a table format
  fix      Scan files in the provided directory recursively and fixes their extensions (when needed). Print report in a table format
  help     Show this help
  version  Print app version and exit

Examples:
  gofixt --dir=~/images scan --exts=jpeg,png,webp
  gofixt --dir=~/files fix
`

	log.Print(usagePrefix)
	flag.PrintDefaults()
}

func validateFlags(dir string, printMode string) {
	if dir == "" {
		log.Fatal("'dir' flag must be provided")
	}
	allowedPrintMode := []string{
		string(internal.PM_ALL),
		string(internal.PM_IMPORTANT),
		string(internal.PM_REPORT),
		string(internal.PM_NONE),
	}
	if !slices.Contains(allowedPrintMode, printMode) {
		log.Fatalf("'print' has invalid value, allowed are: %v", allowedPrintMode)
	}
}
