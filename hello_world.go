package main

import "fmt"

func greetings(name string) string {
	if name == "" {
		return "Hello world"
	}
	return "Hello " + name
}
func main() {
	fmt.Println(greetings("Jhaymes"))
}
