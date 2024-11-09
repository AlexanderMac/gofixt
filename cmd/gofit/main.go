package main

import (
	"flag"
	"log"
	"os"

	"github.com/AlexanderMac/gofit"
)

const VERSION = "0.0.1"

func main() {
	log.SetFlags(0)

	flags := flag.NewFlagSet("gofit", flag.ExitOnError)
	flags.Usage = usage
	dir := flags.String("dir", "", "Scanning directory")

	if err := flags.Parse(os.Args[1:]); err != nil {
		log.Fatalf("Unable to parse args: %v", err)
		return
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
		validateFlags(*dir)
		if err := gofit.Scan(*dir); err != nil {
			log.Fatal(err)
		}
		log.Println("Done!")
	case "fix":
		validateFlags(*dir)
		if err := gofit.Fix(*dir); err != nil {
			log.Fatal(err)
		}
		log.Println("Done!")
	default:
		flags.Usage()
		os.Exit(0)
	}
}

func usage() {
	const usagePrefix = `Usage: gofit [flags] command

Flags:
  --dir Scanning directory (absolute or relative path)

Commands:
  scan     Scans files in the provided directory recursively. Prints files info in a table format
  fix      Scans files in the provided directory recursively and fixes their extensions (when needed). Prints files info in a table format
  help     Shows this help
  version  Prints app version

Examples:
  gofit --dir=~/images scan
  gofit --dir=~/files fix
`

	log.Print(usagePrefix)
	flag.PrintDefaults()
}

func validateFlags(dir string) {
	if dir == "" {
		log.Fatal("'dir' flag must be provided")
	}
}
