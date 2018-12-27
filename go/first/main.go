package main

// Multi import requires ()
import (
	"fmt"
	"math"
	"strconv"

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

// Structs

// Person ()
type Person struct {
	firstName, lastName, city string
	age                       int
}

// Value reciever method - not changing object
func (p Person) greet() string {
	return "Hello " + p.firstName + " " + p.lastName + ", I see you are " + strconv.Itoa(p.age)
}

// Pointer reciever method - changing object
func (p *Person) ageUp() int {
	p.age++
	return p.age
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

	// Pointers
	a := 5
	b := &a
	fmt.Println(a, b)     // 5 0xc0000180e8
	fmt.Printf("%T\n", b) // *int
	// * gets value of pointer
	fmt.Println(*&a) // 5
	fmt.Println(*b)  // 5

	// struct literals
	person1 := Person{firstName: "Grayson", lastName: "Gilmore", age: 36}
	person2 := Person{"Grayson", "Gilmore", "Bellevue", 36}
	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)
	fmt.Println(person1.greet())
	fmt.Println(person2.ageUp())
}
