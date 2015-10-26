package main

import (
	"testing"
)

// at least one argument must be specified
func TestCommandParser(t *testing.T) {
	_, err := parseProgramArgs([]string{""}, false)
	if err == nil {
		t.Fatal("Expected error when passing no arguments")
	}
}

// need -- before the given taskId
func TestMinusMinusMissing(t *testing.T) {
	_, err := parseProgramArgs([]string{"--relengapi-token", "345", "TRZquWniSYmYHlZn_-kLAw"}, false)
	if err == nil {
		t.Fatal("Expected error when not specifying --")
	}
}

// adding unknown parameter causes error
func TestUnknownParameter(t *testing.T) {
	_, err := parseProgramArgs([]string{"--relengapi-token", "345", "--extra-parameter", "--", "TRZquWniSYmYHlZn_-kLAw"}, false)
	if err == nil {
		t.Fatal("Expected error when specifying extra parameter --extra-parameter")
	}
}

// using correct syntax should be ok
func TestValidCommand(t *testing.T) {
	argv := []string{"--relengapi-token", "345", "--", "TRZquWniSYmYHlZn_-kLAw"}
	_, err := parseProgramArgs(argv, false)
	if err != nil {
		t.Fatalf("Expected args %s to be ok, but got error: %s", argv, err)
	}
}

// adding optional parameters works
func TestOptionalParams(t *testing.T) {
	argv := []string{"--relengapi-token", "345", "--port", "678", "--relengapi-host", "pretend-host", "--", "TRZquWniSYmYHlZn_-kLAw"}
	_, err := parseProgramArgs(argv, false)
	if err != nil {
		t.Fatalf("Expected args %s to be ok, but got error: %s", argv, err)
	}
}

// specifying taskId starting with `-` works
func TestLeadingHyphenInTaskId(t *testing.T) {
	argv := []string{"--relengapi-token", "345", "--", "-RZquWniSYmYHlZn_-kLAw"}
	_, err := parseProgramArgs(argv, false)
	if err != nil {
		t.Fatalf("Expected args %s to be ok, but got error: %s", argv, err)
	}
}
