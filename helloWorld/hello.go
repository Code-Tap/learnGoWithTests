package main

import "fmt"

const spanish = "Spanish"
const swahili = "Swahili"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const swahiliHelloPrefix = "Jambo, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {

	switch language {
	case swahili:
		prefix = swahiliHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
