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
	fmt.Println("  countTag - Count the number of tags")
	fmt.Println("  walk - Default usage of gitwalker")
	fmt.Println("  walkByTag - output versions that have tags")
}

func (cli *CLI) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.validateArgs()
	countTagCmd := flag.NewFlagSet("countTag", flag.ExitOnError)
	walkCmd := flag.NewFlagSet("walk", flag.ExitOnError)
	walkByTagCmd := flag.NewFlagSet("walkByTag", flag.ExitOnError)

	walkBare := walkCmd.Bool("bare", false, "do not add commit id to output folder")
	walkByTagBare := walkByTagCmd.Bool("bare", false, "do not add commit id to output folder")

	switch os.Args[1] {
	case "countTag":
		err := countTagCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "walk":
		err := walkCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "walkByTag":
		err := walkByTagCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.printUsage()
	}

	if countTagCmd.Parsed() {
		cli.countTag()
	}

	if walkCmd.Parsed() {
		cli.walk(*walkBare)
	}

	if walkByTagCmd.Parsed() {
		cli.walkByTag(*walkByTagBare)
	}
}
