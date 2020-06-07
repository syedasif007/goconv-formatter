package main

import (
	"fmt"
	"goconv-formatter/formatter"
)

func main() {
	
	// example: convert a float variable into string
	
	varFloat := 10.12

	varString := formatter.StrFloat(varFloat)

	fmt.Println(varFloat, varString)
}