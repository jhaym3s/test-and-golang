package iteration

import "testing"

func TestIteration(t *testing.T) {
	t.Run("alphabets", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})

	t.Run("number", func(t *testing.T) {
		got := Repeat("a")
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	})
	
}

// this is how to run a benchmark function  
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}