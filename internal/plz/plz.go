package plz

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"md0.org/mflag"
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

var (
	client  *pluralize.Client
	choices Opts
	verbose bool
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
		return 0, fmt.Errorf("unknown options: %s", strings.Join(unknowns, ","))
	}
	return b, nil
}

func Main() int {
	cmdname := os.Args[0]
	fs := mflag.NewFlagSet(os.Args[0], mflag.ExitOnError)
	var (
		help   = fs.BoolP("help", "h", false, "print help")
		vrbose = fs.BoolP("verbose", "v", false, "verbose output")
		plur   = fs.BoolP("plural", "p", false, "produce only plural")
		sing   = fs.BoolP("singular", "s", false, "produce only singular")
		opts   = fs.StringP("options", "o", "all", "output options")

		err error
	)
	fs.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: %s [--options/-o ...] <word> [<words...>]\n", cmdname)
		fmt.Fprintf(os.Stderr, "  omit <words> to read stdin\n")
		fmt.Fprintf(os.Stderr, "  --options/-o is a comma-separated list of any combo of:\n")
		fmt.Fprintf(os.Stderr, "    all,isPlural,isSingular,plural,singular\n")
		fmt.Fprintf(os.Stderr, "  --plural/-p, --singular/-s override --verbose and --options\n")
		fmt.Fprintln(os.Stderr)
		fs.PrintDefaults()
	}
	fs.Parse(os.Args[1:])

	if *help {
		fs.Usage()
		return 0
	}

	if *plur && *sing {
		fmt.Fprintln(os.Stderr, "choose one of -plural, -singluar")
		fs.Usage()
		return 1
	}
	if *plur {
		*vrbose = false
		*opts = "plural"
	}
	if *sing {
		*vrbose = false
		*opts = "singular"
	}

	choices, err = parseOpts(*opts)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fs.Usage()
		return 1
	}

	client = pluralize.NewClient()
	verbose = *vrbose

	if fs.NArg() == 0 {
		handleStdin()
	} else {
		handleArgs(fs.Args())
	}
	return 0
}

func handleStdin() {
	fmt.Fprintln(os.Stderr, "ctrl-C to quit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		handleWord(word)
		fmt.Println("--")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func handleArgs(words []string) {
	for i, word := range words {
		handleWord(word)
		if i < len(words)-1 {
			fmt.Println("--")
		}
	}
}

func handleWord(word string) {
	if choices&isPlural == isPlural {
		if verbose {
			fmt.Printf("IsPlural(%s)   => %t\n", word, client.IsPlural(word))
		} else {
			fmt.Println("plur?", client.IsPlural(word))
		}
	}
	if choices&isSingular == isSingular {
		if verbose {
			fmt.Printf("IsSingular(%s) => %t\n", word, client.IsSingular(word))
		} else {
			fmt.Println("sing?", client.IsSingular(word))
		}
	}
	if choices&plural == plural {
		if verbose {
			fmt.Printf("Plural(%s)     => %s\n", word, client.Plural(word))
		} else {
			fmt.Println(client.Plural(word))
		}
	}
	if choices&singular == singular {
		if verbose {
			fmt.Printf("Singular(%s)   => %s\n", word, client.Singular(word))
		} else {
			fmt.Println(client.Singular(word))
		}
	}
}
