package testmod

import "fmt"

// Print prints greetings with salutation and name passed as arguments
func Print(salutation, name string) {
	fmt.Printf("\n%s, %s!\n", salutation, name)
}
