package main

import (
	"fmt"
	"strconv"
)

func SwapNumber(x, y int) (int, int) {
	x = x + y
	y = x - y
	x = x - y
	return x, y
}

func main() {
	x, y := SwapNumber(10, 5)
	fmt.Println("SwapNumber : x = " + strconv.Itoa(x) + " y =" + strconv.Itoa(y))
}
