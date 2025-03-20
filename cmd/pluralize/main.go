package main

import (
	"flag"
	"fmt"

	"github.com/ratrocket/go-pluralize"
	"github.com/ratrocket/go-pluralize/tflags"
)

const (
	appName = "pluralize"
)

func main() {
	var (
		word = flag.String("word", "", "input value")
		cmd  = flag.String("cmd", "All", "command [All|IsPlural|IsSingular|Plural|Singular]")
	)

	flag.Parse()

	if word == nil || len(*word) == 0 {
		fmt.Printf("-word not specified\n")
		return
	}

	pluralize := pluralize.NewClient()

	testCmd := tflags.TestCmdString(*cmd)
	if testCmd.Has(tflags.TestCmdUnknown) {
		fmt.Printf("Unknown -cmd value\nOptions: [All|IsPlural|IsSingular|Plural|Singular]\n")
		return
	}

	if testCmd.Has(tflags.TestCmdIsPlural) {
		fmt.Printf("IsPlural(%s)   => %t\n", *word, pluralize.IsPlural(*word))
	}

	if testCmd.Has(tflags.TestCmdIsSingular) {
		fmt.Printf("IsSingular(%s) => %t\n", *word, pluralize.IsSingular(*word))
	}

	if testCmd.Has(tflags.TestCmdPlural) {
		fmt.Printf("Plural(%s)     => %s\n", *word, pluralize.Plural(*word))
	}

	if testCmd.Has(tflags.TestCmdSingular) {
		fmt.Printf("Singular(%s)   => %s\n", *word, pluralize.Singular(*word))
	}
}
