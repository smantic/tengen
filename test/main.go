package main

import (
	"fmt"
	"os"

	"github.com/smantic/tengen"
)

type Config1 struct {
	Port   int `usage:"port to run your app on"`
	Name   string
	Number int
}

type Config2 struct {
	Config1
	WithMoreFields bool
}

type Config3 struct {
	Config2
	KeepItGoing string
}

type Config4 struct {
	Config1
	// Duplicate field, first one will be set.
	Port int
}

type Config5 struct {
	Field string
	Named struct {
		Field1 string
		Field2 float64
	}
}

func main() {

	if len(os.Args) == 1 {
		fmt.Printf("available commands:\n\ttest subcommand\n\ttest embeded\n\ttest double-embeded\n\ttest duplicate-field\n\ttest nested\n")
		return
	}

	switch os.Args[1] {
	case "subcommand":
		c := Config1{}
		os.Args[1] = "test subcommand"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "embeded":
		c := Config2{}
		os.Args[1] = "test embeded"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "double-embeded":
		c := Config3{}
		os.Args[1] = "test double-embeded"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "duplicate-field":
		c := Config4{}
		os.Setenv("port", "100")
		os.Args[1] = "test duplicate-field"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "nested":
		c := Config5{}
		os.Setenv("field1", "100")
		os.Args[1] = "test nested"
		tengen.Init(&c, os.Args[1:])
		fmt.Printf("%#v\n", c)
	case "help":
		fmt.Printf("helpful message :)")
	default:
	}
}
