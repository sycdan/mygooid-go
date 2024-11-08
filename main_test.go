package main

import (
	"testing"
)

func TestMakeGooid(t *testing.T) {
	tests := []struct {
		input Args
		want  string
	}{
		// neo = 73ef176d9f12809e64363b2b5f4553abecca7aae157327f190323cfa0e42c815
		// the-one = 216c707e6533f5601465b36ac487867a76f98ce43b10141a8538e40c322ec082
		{Args{Name: "Neo", Secret: "The One"}, "89caa2764adf58032739ea0495fd2e9711421ff431953a7140fd7dc5f75b401b"},
		// dan-stace = de51bf63887101b2eb189cfdd447b60cdec1808e10cf3a337070b981b4b58525
		// deranged-hermit = 31def1744b284c8422594ab655121c8ef1c213254ab14c69661feea3cbaf962d
		// general = 0feae16d55365acf07fe9f909834361ba6ee606854746539230bdc84a6a24cee
		// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#%:-_+=? = 942fc2ea554e2e135a6e4b035b73f0d4d173b8b1e0ef3710e2f618a77158643a
		{Args{Name: "Dan Stace", Secret: "Deranged Hermit", Purpose: "General", Characters: "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#%:-_+=?"}, "a457a19d85f629b154eecfd7f8f21c3c127e9db4ba2cd5836eac301aa17b6311"},
	}
	for _, test := range tests {
		got := MakeGooid(test.input)
		if got != test.want {
			t.Errorf("MakeGooid(%s) = %s; want %s", test.input, got, test.want)
		}
	}
}
