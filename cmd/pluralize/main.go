package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"md0.org/pluralize"
)

type Opts uint8

const (
	isPlural Opts = 1 << iota
	isSingular
	plural
	singular

	all = isPlural | isSingular | plural | singular
)

func parseOpts(opts string) (Opts, error) {
	var b Opts
	parts := strings.Split(opts, ",")
	unknowns := []string{}
	for _, part := range parts {
		part = strings.TrimSpace(part)
		switch part {
		case "all":
			b |= all
		case "isPlural":
			b |= isPlural
		case "isSingular":
			b |= isSingular
		case "plural":
			b |= plural
		case "singular":
			b |= singular
		default:
			unknowns = append(unknowns, part)
		}
	}
	if len(unknowns) > 0 {
		return 0, fmt.Errorf("unknown opts: %s", strings.Join(unknowns, "|"))
	}

	return b, nil
}

func main() {
	var (
		cmdname = os.Args[0]
		opts    = flag.String("opts", "all",
			"command [all|isPlural|isSingular|plural|singular], comma sep")
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [-opts ...] <word> [<words...>]\n", cmdname)
		flag.PrintDefaults()
	}
	flag.Parse()

	client := pluralize.NewClient()

	choices, err := parseOpts(*opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(1)
	}

	if flag.NArg() == 0 {
		handleStdin(choices, client)
	} else {
		handleArgs(flag.Args(), choices, client)
	}
}

func handleStdin(choices Opts, client *pluralize.Client) {
	fmt.Println("ctrl-C to end")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		word := strings.TrimSpace(scanner.Text())
		handleWord(word, choices, client)
		fmt.Println("--")
	}
}

func handleArgs(words []string, choices Opts, client *pluralize.Client) {
	for i, word := range words {
		handleWord(word, choices, client)
		if i < len(words)-1 {
			fmt.Println("--")
		}
	}
}

func handleWord(word string, choices Opts, client *pluralize.Client) {
	if choices&isPlural == isPlural {
		fmt.Printf("IsPlural(%s)   => %t\n", word, client.IsPlural(word))
	}
	if choices&isSingular == isSingular {
		fmt.Printf("IsSingular(%s) => %t\n", word, client.IsSingular(word))
	}
	if choices&plural == plural {
		fmt.Printf("Plural(%s)     => %s\n", word, client.Plural(word))
	}
	if choices&singular == singular {
		fmt.Printf("Singular(%s)   => %s\n", word, client.Singular(word))
	}
}
