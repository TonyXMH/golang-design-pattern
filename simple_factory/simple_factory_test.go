package simple_factory

import (
	"testing"
)

func TestNewAPI1(t *testing.T) {
	api := NewAPI(1)
	s := api.Say("Johnson")
	if s != "hi, Johnson" {
		t.Fatal("Type1 test fail")
	}
}
func TestNewAPI2(t *testing.T) {
	api := NewAPI(2)
	s := api.Say("Johnson")
	if s != "hello, Johnson" {
		t.Fatal("Type2 test fail")
	}
}
