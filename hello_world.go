package main


import "fmt"

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


func main() {
	fmt.Println(greetings("Jhaymes", "Spanish"));	
	
}
