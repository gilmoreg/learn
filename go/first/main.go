package main

// Multi import requires ()
import (
	"fmt"
	"math"

	"github.com/gilmoreg/learn/go/first/reverse"
)

// Using var outside of function
var name = "Grayson"

// Cannot use := outside of function

// Basic function
func greeting(str string) string {
	return "Hello " + str
}

// Shortening multiple parameters of same type
func getSum(num1, num2 int) int {
	return num1 + num2
}

func main() {

	// Using := shorthand
	size := 1.3

	// Using const
	const isCool = true
	// isCool = false ERROR

	// Multi assign
	name1, name2 := "Grayson", "Gilmore"

	fmt.Println(greeting(name1), name, isCool, size, name1, name2)

	// %T for type
	fmt.Printf("%T %T\n", name, size)

	// Math
	sq := math.Sqrt(2)
	fmt.Println(sq)

	// Using local module
	r := reverse.Reverse("test")
	fmt.Println(r)

	// Array
	var fruitArr [2]string
	fruitArr[0] = "apple"
	fruitArr[1] = "orange"
	fmt.Println(fruitArr) // [apple orange]

	// Declare and assign
	vegArr := [2]string{"veg1", "veg2"}
	fmt.Println(vegArr) // [veg1 veg2]

	// Slices
	fruitSlice := []string{"grape", "blueberry"}
	fmt.Println(fruitSlice, len(fruitSlice)) // [grape blueberry] 2
	fmt.Println(fruitSlice[1:2])             // [blueberry]

}
