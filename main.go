package main


import (
	"fmt"
	"math/rand"	
)

func main() {
	x := 3;
	const y string = "hi";
	fmt.Println("Hello World");
	fmt.Println(y);
	fmt.Println(x);
	fmt.Println(rand.Intn(10));main_two();
}

func main_two() {
	i := 1;
	for i <= 50 {
		fmt.Println(i);
		i++;

		if i%5 == 0 {
		fmt.Println("floof");
		}

		if i == 50 {break}
	}
	
	//switch statements (shorter way to write if-else statements):
	e := 3
	switch e%5 == 0{
		case true:
			fmt.Println("hi")
		case false:
			fmt.Println("false")

	}

	//fizzbuzz
	h:=0
	for h <= 100 {
		if h%5 == 0 && h%3 == 0 {
			fmt.Println("fizzbuzz")
		}
	switch {
	case h%5 == 0:
		fmt.Println("buzz")
		h++
	case h%3 == 0:
		fmt.Println("fizz")
		h++
	
	} 
	
	
	fmt.Println(h)
	h++;
}
}                   


//Variadic Functions
//variadic functions are