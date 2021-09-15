// declare package
package main

// Importing packages
// the package name is the lasr element of the import path

// factored importing
import (
	"fmt"
	"math"
	"math/rand"
)

// multiple importing - not recommended
// import "fmt"
// import "math"
// import "math/rand"

func add(x int, y int) int {
	return x + y
}

// omit type
func addOmitType(x, y int) int {
	return x + y
}

// function migh return multiple values
func returnMultipleResults(x, y string) (string, string) {
	return x, y
}

// named return values
// naked return
// should be used only in short functions
func returnMultipleResultsNaked(sum int) (x, y int) {
	// variables should be defined on top of the function
	x = sum * 4 / 9
	y = sum - x

	// naked return
	return
}

// Program entry point
func main() {
	fmt.Println("Hello, Wordld!")

	// using the rand package
	// exported names starts with Capital letter - upper case. Lower case names are not exported
	fmt.Println(rand.Intn(10))

	// using the math package
	// exported names starts with Capital letter - upper case. Lower case names are not exported
	fmt.Println(math.Pi)

	fmt.Println(add(10, 5))

	fmt.Println(addOmitType(20, 5))

	// function might return multiple values
	a, b := returnMultipleResults("josé", "almeida")
	fmt.Println(a, b)

	c, d := returnMultipleResultsNaked(17)
	fmt.Println(c, d)
}
