package main

import "fmt"

func main() {
	fmt.Println("Starting the program...")

	safeFunction()

	fmt.Println("Program continues after safeFunction...")
}

func safeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Inside safeFunction, about to cause a panic...")

	panic("Something went wrong!")

	fmt.Println("This line will not execute because of panic.")
}
