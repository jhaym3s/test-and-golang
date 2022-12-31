package injection

import (
	"bytes"
	"testing"
)


func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	 Greet(&buffer, "name")
	 got := buffer.String()
	 want:= "Hello name"

	 if got !=want {
		t.Errorf("got %s wanted %s",got, want)
	 }
	
}