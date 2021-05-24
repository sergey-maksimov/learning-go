//file_size get file size
package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	file := "logo.png"
	fileInfo, err := os.Stat(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())
}
