package main

import "fmt"

func main() {
	fmt.Println("hello world ")
	foo()
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("even numbers only", i)
		}
	}
}
func foo() {
	fmt.Print("foo function was called")
}
