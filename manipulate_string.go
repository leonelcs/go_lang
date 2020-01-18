package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := "        mainha"
	trimGreetings := strings.TrimSpace(greetings)
	helloWorldAndMars := "Hello world and Mars"
	helloWorld := helloWorldAndMars[0:12]
	fmt.Printf("Pega no compasso %s\n", trimGreetings)
	fmt.Println(helloWorld)
	hello2World := "Hello world, world, \"World\""
	newSentence := strings.Replace(hello2World, "world", "mars", 1)
	sentence := "aquele titulo importante do seu texto"
	fmt.Println(newSentence)
	fmt.Println(strings.Title(sentence))
	fmt.Println(strings.ToUpper(sentence))
	fmt.Println(strings.HasSuffix(sentence, "Hello"))
	fmt.Println(strings.HasPrefix(hello2World, "Hello"))

}
