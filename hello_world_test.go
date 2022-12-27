package main

import (
	
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// got := greetings("Jhaymes")
	// want := "Hello Jhaymes"
	// if got != want {
	// 	t.Errorf("got %q want %q", got, want)
	// } 
	t.Run("say hello with name", func(t *testing.T) {
		got := greetings("Jhaymes", "")
		want := "Hello Jhaymes"

		// if got != want {
		// 	t.Errorf("got %q but wanted %q ", got, want)
		// }
		assertCorrectAnswers(t, got, want)
	})

	t.Run("say hello world when there is no name", func(t *testing.T) {
		got := greetings("", "French")
		want := frenchHelloPrefix+"world"

		// if got != want {
		// 	t.Errorf("got %q but wanted %q ", got, want)
		// }
		assertCorrectAnswers(t,got,want)
	})

	//t.Run("")
}


func assertCorrectAnswers(t testing.TB, got, want string)  {
	///testing.TB
	t.Helper()
	///t.Helper() is needed to tell the test suite that this method is a helper
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

///Writing a test is just like writing a function, with a few rules
/// It needs to be in a file with a name like xxx_test.go
/// The test function must start with the word Test
/// The test function takes one argument only t *testing.T
/// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
