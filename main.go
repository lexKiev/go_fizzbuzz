package main

import "fmt"

func main() {
	for i := 0; i <= 100; i++ {
		result := checkValue(i)
		if result == "" {
			fmt.Println(i)
		} else {
			fmt.Println(result)
		}
	}
}
