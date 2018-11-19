package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	const correctOutput = "Hello Testing"
	out := run("echo", Params{correctOutput})
	if out == "hello testing" {
		t.Errorf("Output incorrect, got: %s, want: %s.", string(out), correctOutput)
	}
}
