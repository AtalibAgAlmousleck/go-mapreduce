package main

import "fmt"

// Program to print the day of the week using  switch case
func main() {
	numbers := [5] int {
		11, 22, 33, 44, 55}

	for item := range numbers {
		fmt.Println(numbers[item])
	}

	developer := [3] string {"Ousmane", "Nelson", "Almousleck"}

	for items := range developer {
		fmt.Printf("This is the list of our developers: %v\n", developer[items])
	}

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
}