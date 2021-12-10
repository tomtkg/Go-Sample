package main

import (
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	os.Args = []string{os.Args[0], "../../test/time.json"}
	main()
}
