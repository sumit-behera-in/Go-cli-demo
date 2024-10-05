package main

import (
	"bytes"
	"testing"
)

func Test_greeting(t *testing.T) {

	var buffer = new(bytes.Buffer)

	name := "Sumit"
	greeting(buffer, name)

	expected := "Hello Sumit"
	got := buffer.String()
	if expected != got {
		t.Errorf("expected %v; got %v;", expected, got)
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		main()
	}
}
