package main

//variable and basics

import "fmt"

func main() {
	fmt.Println("rishikeshh")

	name := "rishi"
	age := 20
	fmt.Println(name)
	fmt.Println(age)
	var name2 string
	fmt.Println("enter your name")
	fmt.Scanln(&name2)
	fmt.Println(name2)
	var jink int
	fmt.Println("enter the age of jink")
	fmt.Scanln(&jink)
	fmt.Println(jink)
	if jink > 29 {
		fmt.Println("jink is old")
	} else {
		fmt.Println("jink is young")
	}

	for i := 0; i < jink; i++ {
		fmt.Println(i)
	}
}
