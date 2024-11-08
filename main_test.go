package main

import (
	"testing"
)

func TestMakeGooid(t *testing.T) {
	tests := []struct {
		input Args
		want  string
	}{
		{Args{Name: "Dan Stace", Secret: "Deranged Hermit"}, "686e067590012f005722f0187aa55ae90f6b9eae4d5f2d8074a3a20b3a8bf2ff"},
		{Args{Name: "Neo", Secret: "The One"}, "66da89ca207f8c28ca6086c9d690118229ad7dec1170e5a973f472acb60ac1e9"},
	}
	for _, test := range tests {
		got := MakeGooid(test.input)
		if got != test.want {
			t.Errorf("MakeGooid(%s) = %s; want %s", test.input, got, test.want)
		}
	}
}
