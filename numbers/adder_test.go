package numbers

import (
	//"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("adder function", func(t *testing.T) {
		got := Adder(3, 4)
		want := 7

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})
}

