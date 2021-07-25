package main

import (
	"fmt"
	"os"

	"github.com/smantic/tengen"
)

type Config struct {
	Port   int `usage:"port to run your app on"`
	Name   string
	Number int
}

type Config2 struct {
	Config
	WithMoreFields bool
}

type Config3 struct {
	Config2
	KeepItGoing string
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("available commands:\n\ttest sub\n\ttest command\n\ttest list\n")
		return
	}

	switch os.Args[1] {
	case "sub":
		c := Config{}
		os.Args[1] = "test sub"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "command":
		c := Config2{}
		os.Args[1] = "test command"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "list":
		c := Config3{}
		os.Args[1] = "testlist"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "help":
		fmt.Printf("helpful message :)")
	default:
	}
}
