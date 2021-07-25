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
func Init(c interface{}, args []string) {

	flags := extract(os.Args[0], c)
	err := flags.Parse(args[1:])
	if err != nil {
		log.Printf("failed to set flags", err)
		flags.Usage()
	}
}

func extractEnv(s reflect.Value) {

	typ := s.Type()
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		value := s.FieldByName(field.Name)

		if value.Kind() == reflect.Ptr {
			value = value.Elem()
		}

		switch value.Kind() {
		}
	}
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
		case *interface{}:
			//extractIntoFlagSet(&(s.Field(i).Interface{}), flags)
		default:
			continue
		}
	}
}
