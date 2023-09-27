package main

import "fmt"

func makeMult(base int) func(int) int {
	return func(factor int) int { return base * factor}
}

func main(){
	multBy2 := makeMult(2)
	multBy3 := makeMult(3)

	for i := 0; i < 5; i++ {
		result :=fmt.Sprintf("%v * 2 = %v, %v * 3 = %v", i, multBy2(i), i, multBy3(i))
		fmt.Println(result)
	}
}