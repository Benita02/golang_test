package main

import "fmt"

func main() {
	fmt.Println("Benita is soooo cool")
	var name string = "George of the Jungle"
	fmt.Println(name)
	name2 := "Rohit"
	fmt.Println(name2)
	name2 = "Da new name ain't Rohit but Benita"

	var num int = 3
	fmt.Println(num)

	num = 76
	fmt.Println(num)

	var flag bool = true
	fmt.Println(flag)

	var pie float32 = 3.14
	fmt.Println(pie)

	var age int = -89
	// if age > 0 && age <= 2 {
	// 	fmt.Println("You are an infant")
	// } else if age > 2 && age <= 5 {
	// 	fmt.Println("You are a toddler")
	// } else if age > 5 && age <= 12 {
	// 	fmt.Println("You are a child")
	// } else if age > 12 && age <= 18 {
	// 	fmt.Println("You are a teenager, you stubborn brat")
	// } else if age > 18 && age <= 35 {
	// 	fmt.Println("You're a bloody adult, go get a job")
	// } else if age > 35 && age <= 49 {
	// 	fmt.Println("You're middle aged, hope you made the right choices")
	// } else if age > 50 {
	// 	fmt.Println("You're an elder, congrats, this wicked world didn't kill you......YET!")
	// } else {
	// 	fmt.Println("Invalid Input, try your shit again, dumbass!")
	// }

	switch {
	case age > 0 && age <= 2:
		fmt.Println("You are an infant")
	case age > 2 && age <= 5:
		fmt.Println("You are a toddler")
	case age > 5 && age <= 12:
		fmt.Println("You are a child")
	case age > 12 && age <= 18:
		fmt.Println("You are a teenager, you stubborn brat")
	case age > 18 && age <= 35:
		fmt.Println("You're a bloody adult, go get a job")
	case age > 35 && age <= 49:
		fmt.Println("You're middle aged, hope you made the right choices")
	case age > 50:
		fmt.Println("You're an elder, congrats, this wicked world didn't kill you......YET!")
	default:
		fmt.Println("Invalid Input, try your shit again, dumbass!")

	}
}
