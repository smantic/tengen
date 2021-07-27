package tengen

import (
	"flag"
	"log"
	"os"
	"reflect"
	"strings"
)

// Init expects a pointer to a config so we can fill the values using reflection.
// args[0] is name of command (recomend including subcommand like: "git" + os.Args[1] -> git help)
func Init(c interface{}, args []string) flag.FlagSet {

	if len(args) == 0 {
		log.Fatalln("expected args with application name.")
	}

	flags := extract(args[0], c)
	err := flags.Parse(args[1:])
	if err != nil {
		flags.Usage()
	}
	flags.VisitAll(extractEnv)
	return flags
}

func extract(name string, c interface{}) flag.FlagSet {

	if reflect.ValueOf(c).Kind() != reflect.Ptr {
		log.Fatalf("expected a pointer to a struct for reflection")
	}

	flags := flag.NewFlagSet(name, flag.ExitOnError)
	extractIntoFlagSet(c, flags)

	return *flags
}

func extractIntoFlagSet(c interface{}, flags *flag.FlagSet) {
	val := reflect.ValueOf(c).Elem()
	typ := val.Type()

	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		addr := val.Field(i).Addr().Interface()
		name := strings.ToLower(field.Name)
		usage := field.Tag.Get("usage")

		if flag := flags.Lookup(name); flag != nil {
			// if we try to set two flags with the same name it will panic
			continue
		}

		switch ptr := addr.(type) {
		case *int:
			flags.IntVar(ptr, name, *ptr, usage)
		case *uint:
			flags.UintVar(ptr, name, *ptr, usage)
		case *float64:
			flags.Float64Var(ptr, name, *ptr, usage)
		case *string:
			flags.StringVar(ptr, name, *ptr, usage)
		case *bool:
			flags.BoolVar(ptr, name, *ptr, usage)
		case interface{}:
			extractIntoFlagSet(ptr, flags)
		default:
			continue
		}
	}
}

// get flag value from env
func extractEnv(f *flag.Flag) {
	v := os.Getenv(f.Name)
	if v == "" {
		return
	}

	f.Value.Set(v)
}
