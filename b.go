package main

import "fmt"

func main() {
	fmt.Println("first go program")
	name := "rishikesh"
	age := 3
	fmt.Println(name)
	fmt.Println(age)

	if age > 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("minor")
	}
  
}
