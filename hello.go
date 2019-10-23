package main

import (
		"fmt"
		"math"
		"math/rand"
		"github.com/Kunal/stringutil"
		)

func main() {
	var c = 3
	var d = 3*c
	fmt.Printf("hello, world\n")
	fmt.Println(stringutil.Reverse("!oG ,olleH"))
	rand.Seed(11)
	fmt.Println("this is a random number", rand.Intn(10))
	fmt.Println("this is a random number", math.Sqrt(49))
	fmt.Println(pi)
	fmt.Println(math.Pi)
	sum := func(a, b int) int { return a+b } (3, 4)
	fmt.Println(sum)
	d:=5
	fmt.Println(simpleInterest(1000.00, 9.0, 12))
	a, b:=stringutil.AddMultiple(16,4)
	fmt.Println(a,b)
	fmt.Println(c,d)
}

func pi() float64	{
	return math.Pi
}

func simpleInterest(principal float64, rate float64, duration float64) float64{
	return (principal * rate * duration)
}

