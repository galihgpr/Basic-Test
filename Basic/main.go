package main

import "fmt"

func main() {
	val1 := "Frontend"
	val2 := "Backend"
	length := 50
	num1 := 3
	num2 := 5
	total := 0
	for i := 1; i <= length; i++ {
		if i%num1 == 0 && i%num2 == 0 {
			total++
		}
	}
	counter := 0

	for i := 1; i <= length; i++ {
		coma := ""
		if i == length {
			coma = ""
		} else {
			coma = ","
		}
		if i%num1 == 0 && i%num2 == 0 {
			counter++
			if counter == total {
				fmt.Print(val1 + " " + val2 + coma)
			} else {
				fmt.Println(val1)
				fmt.Print(val2 + coma)
			}
		} else if i%num1 == 0 {
			fmt.Print(val1 + coma)
		} else if i%num2 == 0 {
			fmt.Print(val2 + coma)
		} else {
			fmt.Print(i, coma)
		}
		if i == 50 {
			fmt.Println()
		}
	}
}
