package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "hello czq hello"
	var s2 string = "hello czq hello"
	fmt.Println(!strings.EqualFold(s, s2))
}

