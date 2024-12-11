package packageB

import "fmt"

func decrement(i int) int {
	return i - 1
}

func PrintDecrement(i int) {
	fmt.Println("Decremented value:", decrement(i))
}
