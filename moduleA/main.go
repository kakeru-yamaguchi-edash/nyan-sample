package main

import (
	"moduleA/packageA"
	"moduleA/packageB"
)

func main() {
	packageA.PrintIncrement(1) // result: "Incremented value: 2"
	packageB.PrintDecrement(1) // result: "Decremented value: 0"
}
