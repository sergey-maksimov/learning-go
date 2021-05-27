// The game generates a number from 1 to 100 and asks the user to guess it
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Get the current date and time in integer format
	seconds := time.Now().Unix()
	// Random number generator function
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я выбрал случайное число от 1 до 100")
	fmt.Println("Сможешь угадать его?")
	// Showing the generated number
	fmt.Println(target)
}
