package hello

import (
	"fmt"
	"log"

	"github.com/coghost/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
    names := []string {"Gladys", "Samantha", "Darrin"}
	message, err := greetings.HelloGuys(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
