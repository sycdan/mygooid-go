package main

import (
	"testing"
)

func TestMakeGooid(t *testing.T) {
	tests := []struct {
		input Args
		want  string
	}{
		// dan-stace = de51bf63887101b2eb189cfdd447b60cdec1808e10cf3a337070b981b4b58525
		// deranged-hermit = 31def1744b284c8422594ab655121c8ef1c213254ab14c69661feea3cbaf962d
		// general = 0feae16d55365acf07fe9f909834361ba6ee606854746539230bdc84a6a24cee
		// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#%:-_+=? = 942fc2ea554e2e135a6e4b035b73f0d4d173b8b1e0ef3710e2f618a77158643a
		{Args{Name: "Dan Stace", Secret: "Deranged Hermit"}, "a457a19d85f629b154eecfd7f8f21c3c127e9db4ba2cd5836eac301aa17b6311"},
		{Args{Name: "Neo", Secret: "The One"}, "66da89ca207f8c28ca6086c9d690118229ad7dec1170e5a973f472acb60ac1e9"},
	}
	for _, test := range tests {
		got := MakeGooid(test.input)
		if got != test.want {
			t.Errorf("MakeGooid(%s) = %s; want %s", test.input, got, test.want)
		}
	}
}
