package main

import (
	"testing"
)

func TestMakeGooid(t *testing.T) {
	tests := []struct {
		input Args
		want  string
	}{
		{Args{Name: "Dan Stace", Secret: "Deranged Hermit"}, "ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3"},
		{Args{Name: "Neo", Secret: "The One"}, "a767aaaab7bfe7a652a2cd4076d443de1591315cd33462375ca5a3a9067df41d"},
	}
	for _, test := range tests {
		got := MakeGooid(test.input)
		if got != test.want {
			t.Errorf("MakeGooid(%s) = %s; want %s", test.input, got, test.want)
		}
	}
}
