package slice

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("test sum for a fixed size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		assertCorrectAnswers(t, got, want)
	})

	t.Run("test sum for a various size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		assertCorrectAnswers(t, got, want)

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9 })
	want := []int{3, 15}

	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }
	///invalid operation: cannot compare got != want (slice can only be compared to nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v ", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}

func assertCorrectAnswers(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
