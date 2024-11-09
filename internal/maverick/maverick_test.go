package maverick

import (
	"testing"

	"github.com/sycdan/mygooid-go/internal/renji"
	"github.com/sycdan/mygooid-go/internal/utils"
)

func TestNext(t *testing.T) {
	const SEED_HASH = "ce568acbe94abae40f3c9d814fd82dbdb6eac76157f2dec42837a471a73e74f3"
	tests := []struct {
		input []interface{}
		want  interface{}
	}{
		{[]interface{}{"a", "b", "c"}, "c"},
		{[]interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2},
	}
	for _, test := range tests {
		var rng = renji.NewRenji(SEED_HASH)
		shuffler := NewMaverick(test.input, rng)
		got := shuffler.Next()
		if got != test.want {
			t.Errorf("NewMaverick(%s).Next() = %s; want %s", test.input, utils.ToString(got), utils.ToString(test.want))
		}
	}
}
