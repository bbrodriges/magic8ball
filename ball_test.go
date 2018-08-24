package main

import (
	"plugin"
	"testing"
)

func TestPlugin(t *testing.T) {
	plug, err := plugin.Open("./ball-darwin-amd64.so")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	symb, err := plug.Lookup("Answer")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	answerer, ok := symb.(func(string) string)
	if !ok {
		t.Fatalf("Expecting symbol type to be `func(string)string`")
	}

	_ = answerer("Will this project have more than 100 stars on Github?")
}
