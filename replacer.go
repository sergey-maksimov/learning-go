package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "I l0ve G0lang!"
	replacer := strings.NewReplacer("0", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)
}
