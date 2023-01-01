package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const englishHelloPrefix = "Hello "
const spanishHelloPrefix = "Hola "
const frenchHelloPrefix = "Bonjour "

 


func getGreetingPrefix(language string) (prefix string){
	switch language {
	case "French":
		prefix = frenchHelloPrefix
	case "Spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
		
	}
	return prefix
}

func greetings(name, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(language) + name
}


func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "James")
}

///Run the program and go to http://localhost:5001. You'll see your greeting function being used.


func main() {
	fmt.Println(greetings("Jhaymes", "Spanish"));	
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func (u string)  {
			results[u] = wc(u)
		}(url)
		time.Sleep(2 * time.Second)
	}

	return results
}