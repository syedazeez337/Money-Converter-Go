package main

import (
	"fmt"
	"strconv"
)

func main() {
	value, err := strconv.ParseInt("1010", 2, 64)
	fmt.Printf("%v - %v\n", value, err)
}
