package main

import "fmt"

func main() {

	f, b := 60, "668"
	fmt.Println(f)
	fmt.Println(b)
	pick(5, 4)
	pick(2, 2)
}

func pick(a, b int) {
	if (a == b) {
		fmt.Println("Equals") 
	} else {
		fmt.Println("Sorry")
	}
}

func max(a, b int) int {
	if (a>b){
		return a
	} else {
		return b
	}
}
