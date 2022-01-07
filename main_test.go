// Copyright (c) 2022, Microsoft Corporation, Sean Hinchee
// Licensed under the MIT License.
package main

import "testing"

func TestParse(t *testing.T) {
	expected := `U+222B '∫'`
	out := parse("\u222B")
	if out != expected {
		t.Error("Literal \\u1234 failed, got:", out)
	}

	out = parse("\\u222B")
	if out != expected {
		t.Error("\\\\u1234 failed, got:", out)
	}

	out = parse("u222B")
	if out != expected {
		t.Error("u1234 failed, got:", out)
	}

	out = parse("U+222B")
	if out != expected {
		t.Error("U+1234 failed, got:", out)
	}
}
