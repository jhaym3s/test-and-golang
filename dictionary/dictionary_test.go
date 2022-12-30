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
		want := ErrNotFound

		if err != want {
			t.Errorf("got %q wanted %q ", err, want)
		}
		if err == nil {
			t.Fatalf("this was never suppose to happen")
		}
	})
	
}

func TestAdd(t *testing.T) {

	t.Run("new word",func(t *testing.T) {
		dictionary := Dictionary{}
	word := "test"
	definition := "this is just a test"

	err := dictionary.Add(word, definition)

	assertError(t,err,nil)

	assertDefinition(t, dictionary, word, definition)

	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
	
}

func TestUpdate(t *testing.T) {
	dictionary := Dictionary{"word": "meaning"}
	t.Run("existing word", func(t *testing.T) {
		got := dictionary.Update("word","new meaning")
		if got != nil{
		t.Errorf("got %q", got)
		}
	})

	t.Run("new words", func(t *testing.T) {
		got := dictionary.Update("no word", "meaning")
		want := ErrNotFound
		if got != want {
			t.Errorf("got %q wanted %q", got,want)
		}
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dictionary := Dictionary{word: "test definition"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)
	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", word)
	}
}


func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertError(t testing.TB,got,want error)  {
	t.Helper()
	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}