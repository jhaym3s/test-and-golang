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
	got := SumAll([]int{1, 2}, []int{0, 9, 5,1})
	// So turns out that is one slice has more elements that it 
	//will continue to add to last element other slides have in common
	want := []int{3, 15}

	// if got != want {
	// 	t.Errorf("got %v want %v", got, want)
	// }
	///invalid operation: cannot compare got != want (slice can only be compared to nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v ", got, want)
	}
}

func assertCorrectAnswers(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
