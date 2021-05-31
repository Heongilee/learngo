package main

import (
	"fmt"
)

// func lenAndUpper(name string) (int, string) {
// 	return len(name), strings.ToUpper(name)
// }

func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	// totalLength, upperName := lenAndUpper("heongi")

	repeatMe("nico", "heongi", "dal", "marl", "flynn")
}
