package main

func checkValue(i int) (value string) {

	if i%3 == 0 {
		value += "Fizz"
	}
	if i%5 == 0 {
		value += "Buzz"
	}

	return
}
