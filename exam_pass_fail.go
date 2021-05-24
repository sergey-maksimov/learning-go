//exam_pass_fail did the student pass the exam
package main

//importing packages
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//requesting data from the user
	fmt.Print("Enter a grade: ")
	//getting text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	//return the text entered by the user
	input, err := reader.ReadString('\n')
	//if the error is not equal to nil
	if err != nil {
		//report an error and abort the program
		log.Fatal(err)
	}
	//removing the newline character from the input string
	input = strings.TrimSpace(input)
	//converting a string to a value float64
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	//declaring a variable
	var status string
	//if the value is equal to or greater than 60 the variable is assigned a string "passing"
	if grade >= 60 {
		status = "passing"
	//otherwise, it is assigned "failing"
	} else {
		status = "failing"
	}
	//output the value and the result
	fmt.Println("A grade of", grade, "is", status)
}
