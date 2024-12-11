package packageA

import (
	"fmt"
)

func callFunc1(i int) int { // パッケージ外には非公開
	return increment(i)
}

func PrintIncrement(i int) {
	fmt.Println("Incremented value:", callFunc1(i))
}
