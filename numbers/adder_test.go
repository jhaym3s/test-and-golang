package numbers

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	t.Run("first", func(t *testing.T) {
		got := Adder(3, 4)
		want := 7

		if got != want {
			t.Errorf("got %q expected %q", got, want)
		}
	})
}

func ExampleAdder()  {
	sum := Adder(1, 5)
	fmt.Println(sum)
}
