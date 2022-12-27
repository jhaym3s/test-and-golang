package main

import "fmt"

const englishHelloPrefix = "Hello "

func greetings(name string) string {
	if name == "" {
		return "Hello world"
	}
	return englishHelloPrefix + name
}
func main() {
	fmt.Println(greetings("Jhaymes"))
}
