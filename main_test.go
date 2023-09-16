package main

import (
	"fmt"
	"testing"
)

func TestParseL(t *testing.T) {
	tests := map[string]string{
		"1":             "1",
		"1 + 2":         "{1 + 2}",
		"1 + 2 * 3":     "{{1 + 2} * 3}",
		"1 * 2 + 3":     "{{1 * 2} + 3}",
		"1 * 2 + 3 * 4": "{{{1 * 2} + 3} * 4}",
		"1 + 2 * 3 + 4": "{{{1 + 2} * 3} + 4}",
	}
	for in, want := range tests {
		tokens := newTokens(in)
		expr := parseL(tokens)
		got := fmt.Sprintf("%v", expr)
		assertEqual(t, want, got)
	}
}

func TestParseR(t *testing.T) {
	tests := map[string]string{
		"1":             "1",
		"1 + 2":         "{1 + 2}",
		"1 + 2 * 3":     "{1 + {2 * 3}}",
		"1 * 2 + 3":     "{1 * {2 + 3}}",
		"1 * 2 + 3 * 4": "{1 * {2 + {3 * 4}}}",
		"1 + 2 * 3 + 4": "{1 + {2 * {3 + 4}}}",
	}
	for in, want := range tests {
		tokens := newTokens(in)
		expr := parseR(tokens)
		got := fmt.Sprintf("%v", expr)
		assertEqual(t, want, got)
	}
}

func TestParse(t *testing.T) {
	tests := map[string]string{
		"1":             "1",
		"1 + 2":         "{1 + 2}",
		"1 + 2 * 3":     "{1 + {2 * 3}}",
		"1 * 2 + 3":     "{{1 * 2} + 3}",
		"1 * 2 + 3 * 4": "{{1 * 2} + {3 * 4}}",
		"1 + 2 * 3 + 4": "{{1 + {2 * 3}} + 4}",
	}
	for in, want := range tests {
		tokens := newTokens(in)
		expr := parse(tokens, lowest)
		got := fmt.Sprintf("%v", expr)
		assertEqual(t, want, got)
	}
}

func assertEqual(t *testing.T, want, got string) {
	t.Helper()
	if want == got {
		return
	}
	t.Fatalf("want != got\n%s\n%s", want, got)
}
