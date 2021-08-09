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

	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == swahili {
		return swahiliHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
