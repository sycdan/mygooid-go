package maverick

import (
	"testing"

	"github.com/sycdan/mygooid-go/internal/renji"
	"github.com/sycdan/mygooid-go/internal/utils"
)

const SEED_HASH = "ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3"

var rng = renji.NewRenji(SEED_HASH)

func TestNext(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"a", "b", "c"}, "c"},
	}
	for _, test := range tests {
		shuffler := NewMaverick(test.input, rng)
		got := shuffler.Next()
		if got != test.want {
			t.Errorf("NewMaverick(%s).Next() = %s; want %s", test.input, utils.ToString(got), utils.ToString(test.want))
		}
	}
}
