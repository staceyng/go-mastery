// Instructions:
// Complete each task below without looking at references.
// Try to recall the syntax from memory and write valid Go code.
package main

import (
	"fmt"
)

// 1. Declare a package-level variable.
var defaultNum int

// 2. Declare a variable `name` of type string and assign it your name.
var name string
name = "stacey"

// 3. Declare an integer variable `age` and assign it a value.
var age int = 10

// 4. Declare a float variable `height` and assign it a value.
var height float = 1.6

// 5. Declare a boolean variable `isCodingFun` and assign it `true`.
var isCodingFun bool = true

// 6. Declare multiple variables in a single line `city`, `country`, `population`.
var city, country, population string

// 7. Use the shorthand syntax to declare and initialize a variable for your favorite language.
language := "golang"

// 8. Declare a constant `pi` with a value of 3.14159.
const pi float = 3.14159

// 9. Swap two variables without using a temporary variable.
var a, b = 5, 10
a, b = b, a

// 10. Print the **type** of a variable using the golang formatting lib
fmt.Printf("Type of pi variable is %T \n", pi)

// 11. Declare a zero-valued variable without assigning anything. Print its default value.
var defaultText string
fmt.Printf("Default value of defaultText: %s \n", defaultText)

// 12. Declare a pointer variable that points to an integer and modify the value using the pointer.
var pointerVar *int
val := 4
pointerVar = &val
fmt.Printf("New value of pointerVar: %d \n", *pointerVar)

// 13. Create a function that returns multiple values and store them in variables.
func getPersonalDetails() (string, string, int) {
	firstname := "stacey"
	lastname := "ng"
	age := 10
	return firstname, lastname, age
}
fn, ln, age := getPersonalDetails()

// 14. Convert an integer to a float and vice versa.
var someInt int 4
var someFloat float32 = 3.5

convertedFloatToInt := int(someFloat)
convertedIntToFloat := float32(someInt)

fmt.Printf("New value of convertedInt: %v \n", convertedIntToFloat)
fmt.Printf("New value of convertedFloat: %v \n", convertedFloatToInt)

// 15. Use `const` with an untyped constant and assign it to different typed variables.
const pi = 3.14159
var radius int = 5
var area float64 = pi * float64(radius) * float64(radius)
fmt.Printf("\nArea: %v\n", area)

// 16. Declare and use a shadowed variable inside a function.
func shadowName() {
    name := "John"
    fmt.Printf("Outer name: %s\n", name)

    if true {
        name := "Jane"  // Shadows outer name
        fmt.Printf("Inner name: %s\n", name)
    }
    
    fmt.Printf("Outer name after if block: %s\n", name)  // Still "John"
}

// 17. Use the new keyword to allocate memory for an integer and modify its value.
newIntPointer := new(int)
fmt.Printf("Initial pointer value: %d\n", *ptr)
fmt.Printf("Pointer address: %p\n", ptr)

*newIntPointer = 99
fmt.Printf("Modified value: %d\n", *ptr)

// 18. Use shorthand assignment inside an `if` statement.
func example() {
	if x := 10; x > 5 {
		fmt.Printf("x is %d and greater than 5\n", x)
	}
}

// 19. Use a function closure to modify an outer variable.
func createCounter(start int) func() int {
    count := start
    
    return func() int {
        current := count
        count++
        return current
    }
}
createCounter(10)