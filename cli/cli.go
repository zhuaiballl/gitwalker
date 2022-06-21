package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  walk - Default usage of gitwalker")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	walkCmd := flag.NewFlagSet("walk", flag.ExitOnError)

	switch os.Args[1] {
	case "walk":
		err := walkCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		err := walkCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	}

	if walkCmd.Parsed() {
		cli.walk()
	}
}
