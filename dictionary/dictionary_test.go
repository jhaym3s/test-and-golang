package dictionary

import "testing"

func TestSearch(t *testing.T) {
 dictionary := Dictionary{"test": "test meaning"}
	t.Run("known words", func(t *testing.T) {
		got,_ := dictionary.Search("test")
		want := "test meaning"

		if got != want {
			t.Errorf("got %q wanted %q ", got, want)
		}
	})

	t.Run("unknown words", func(t *testing.T) {
		_,err := dictionary.Search("unknown")
		want := errNotFound

		if err != want {
			t.Errorf("got %q wanted %q ", err, want)
		}
		if err == nil {
			t.Fatalf("this was never suppose to happen")
		}
	})
	
}