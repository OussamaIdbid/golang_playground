package main

import "fmt"

func main() {

	var names = []string{"string", "string2"}

	for i, value := range names {
		fmt.Println(i, value)
	}
}
