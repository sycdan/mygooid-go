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
		// <blank> = e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
		{Args{Name: "Neo", Secret: "The One"}, "38bcfb2bb51a3c125e00f59fd1bca2cbc99132d336bbe3e189311cf90a77f965"},
		// dan-stace = de51bf63887101b2eb189cfdd447b60cdec1808e10cf3a337070b981b4b58525
		// deranged-hermit = 31def1744b284c8422594ab655121c8ef1c213254ab14c69661feea3cbaf962d
		// general = 0feae16d55365acf07fe9f909834361ba6ee606854746539230bdc84a6a24cee
		{Args{Name: "Dan Stace", Secret: "Deranged Hermit", Purpose: "General", Characters: "ignored"}, "77281e96895dc2364b22e7192aec81e3bb8d5c265ec4aedc36cfebfc7e598f48"},
	}
	for _, test := range tests {
		got := MakeGooid(test.input)
		if got != test.want {
			t.Errorf("MakeGooid(%s) = %s; want %s", test.input, got, test.want)
		}
	}
}
