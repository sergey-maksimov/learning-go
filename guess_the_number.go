// The game generates a number from 1 to 100 and asks the user to guess it
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	// Сreating a bufio. NewReader for reading keyboard input
	reader := bufio.NewReader(os.Stdin)
	// Request a number from the user
	fmt.Print("Введите число:")
	// Read the data entered by the user before pressing enter
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// Deleting a newline character
	input = strings.TrimSpace(input)
	// Convert to an integer
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if guess < target {
		fmt.Println("Ваше предположение низкое")
	} else if guess > target {
		fmt.Println("Ваше предположение высокое")
	}
}
