package sushigo

import "testing"

func TestValidateCards(t *testing.T) {
	si := SimulationInput{
		1,
		[]string{"a", "b", "c", "a", "a"},
	}
	if si.ValidateCards() {
		t.Error("Gave invalid cards and expected failure but received succes.")
	}
}

func TestValidateCards2(t *testing.T) {
	si := SimulationInput{
		1,
		[]string{"maki-1", "salmon", "pudding", "pudding"},
	}
	if !si.ValidateCards() {
		t.Error("Found invalid card.")
	}
}
