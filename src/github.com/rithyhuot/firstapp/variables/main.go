package main

import (
	"fmt"
	"strconv"
)

// variables in Go has to be used

/*
 * lowercase variable scopes to the package level - var test
 * uppercase variable scopes to the global level and is exported - var Test
 */

/*
 * Rule of thumb for determining variable name in Go
 *  - the length of the variable name is derived by the life time use of the variable
 */

/* when setting variable outside a function, you can't use the
 * shorthand variable declaration - (:=) i.e. k := 99
 */

/*
 * var i int = 24
 * var j string = "hello"
 * var k int = 25
 * var t string = "world"
 *
 * Same as below
 */

var (
	i int    = 24
	j string = "hello"
	k int    = 25
	t string = "world"
)

func main() {
	// use when you need to don't need to initialize a value
	// var i int
	// i = 42

	/* use when you want more control over the type. Such as wanting the
	 * to be a float vs int
	 */
	// var j float32 = 27

	// let Go determines the variable type
	// k := 99

	/*
	 * Convert between types
	 * var i int = 5
	 * var j float32 = float32(i)
	 */

	var i = 42
	var j string
	j = strconv.Itoa(i)

	fmt.Printf("%v, %T", j, j)
}
