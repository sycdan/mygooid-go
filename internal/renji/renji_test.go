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
		got := New(test.input).Next()
		if got != test.want {
			t.Errorf("New(%s).Next() = %s; want %s", test.input, utils.ToString(got), utils.ToString(test.want))
		}
	}
}
