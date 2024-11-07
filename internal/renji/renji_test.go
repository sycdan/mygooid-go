package renji

import (
	"testing"

	"github.com/sycdan/mygooid-go/internal/utils"
)

func TestNextProducesExpectedFloatForGivenHash(t *testing.T) {
	tests := []struct {
		input string
		want  float64
	}{
		{"ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3", 0.6555598267699682},
		{"a767aaaab7bfe7a652a2cd4076d443de1591315cd33462375ca5a3a9067df41d", 0.40785707194054543},
	}
	for _, test := range tests {
		got := NewRenji(test.input).Float64()
		if got != test.want {
			t.Errorf("NewRenji(%s).Float64() = %s; want %s", test.input, utils.ToString(got), utils.ToString(test.want))
		}
	}
}

func TestNextBetween(t *testing.T) {
	tests := []struct {
		hash string
		min  int
		max  int
		want int
	}{
		{"ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3", 0, 1, 1},
		{"ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3", 1, 1, 1},
		{"a767aaaab7bfe7a652a2cd4076d443de1591315cd33462375ca5a3a9067df41d", 1, 2, 1},
		{"a767aaaab7bfe7a652a2cd4076d443de1591315cd33462375ca5a3a9067df41d", 3, 9, 5},
	}
	for _, test := range tests {
		rng := NewRenji(test.hash)
		// Generate a bunch to make sure they are within the boundaries, but check the first is consistent.
		for i := 0; i < 1000; i++ {
			got := rng.NextBetween(test.min, test.max)
			if got < test.min || got > test.max {
				t.Errorf("NewRenji(%s).NextBetween() = %d; want within [%d, %d]", test.hash, got, test.min, test.max)
			}
			if i == 0 && got != test.want {
				t.Errorf("NewRenji(%s).NextBetween() = %d; want %d", test.hash, got, test.want)
			}
		}
	}
}
