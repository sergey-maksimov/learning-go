package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(reflect.TypeOf(3))
	fmt.Println(reflect.TypeOf("Hello"))
	fmt.Println(reflect.TypeOf(10.5))
	fmt.Println(reflect.TypeOf(true))
}
