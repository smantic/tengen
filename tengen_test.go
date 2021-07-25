package tengen

import (
	"os"
	"strconv"
	"testing"
)

type Config struct {
	Port   string `usage:"port to run your app on"`
	Name   string
	Number int
}

func TestApp(t *testing.T) {

	os.Setenv("Port", ":111")
	os.Setenv("Number", "111")

	c := Config{}
	Init(&c)

	if c.Port != os.Getenv("Port") {
		t.Errorf("expected port env var to be: %s. Got: %s\n", os.Getenv("Port"), c.Port)
	}

	numFromEnv, _ := strconv.Atoi(os.Getenv("Number"))
	if c.Number != numFromEnv {
		t.Errorf("expected port env var to be: %d. Got: %d\n", numFromEnv, c.Number)
	}
}
