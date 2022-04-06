package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Your Greeting: ")
	phrase, _ := reader.ReadString('\n')
	phrase = strings.TrimSpace(phrase)

	fmt.Print("Your Name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	message := Greeting(name, phrase)
	fmt.Println(message)

}
func Greeting(name, message string) (salutation string) {
	salutation = fmt.Sprintf("%s, %s", message, name)
	return
}
