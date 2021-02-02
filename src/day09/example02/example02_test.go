package exampel02

import "testing"

func TestFrTranslation(t *testing.T) {
	got := translate("fr")
	want := "Bonjour "
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}

func TestUSTranslation(t *testing.T) {
	got := translate("en-US")
	want := "Hello "
	if got != want {
		t.Fatalf("Expected %q, got %q", want, got)
	}
}
