package main

import "fmt"

var _global_ int = 9

func main() {
	var i, j int = 1, 2

	// short
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java, _global_)
}
