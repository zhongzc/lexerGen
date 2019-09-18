package reparser

import "testing"

func TestNormal(t *testing.T) {
	r, err := Parse("abc*((cat)*|dog)")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "ab(c)*((cat)*|dog)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestCharset(t *testing.T) {
	r, err := Parse("[a-c1-4_]")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "(a|b|c|1|2|3|4|_)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}

func TestEscCharset(t *testing.T) {
	r, err := Parse("[a-c\\]1-4_\\-]")
	if err != nil {
		t.Fatalf("reparser.Parse() failed: %s", err.Error())
	}
	expected := "(a|b|c|]|1|2|3|4|_|-)"
	if r.REString() != expected {
		t.Fatalf("reparser.Parse() expected %q. got %q", expected, r.REString())
	}
}