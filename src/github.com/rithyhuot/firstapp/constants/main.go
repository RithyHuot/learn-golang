package main

import (
	"fmt"
)

const (
	a = iota // scope to this constant block
	b
	c
)

const (
	// errorSpecialist = iota
	// _ = iota // generate a value but throw it away
	_ = iota + 5
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	_  = iota             // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota) // 2^10
	MB                    // 2^100
	GB                    // 2^1000
	TB                    // 2 ^10000
	PB                    // 2 ^100000
	EB                    // 2 ^1000000
	ZB                    // 2 ^10000000
	YB                    // 2 ^100000000
)

const (
	isAdmin             = 1 << iota // 2^0 = 1
	isHeadquarters                  // 2^1 = 2
	canSeeFinancials                // 2^2 = 4
	canSeeAfrica                    // 2^3 = 8
	canSeeAsia                      // 2^4 = 16
	canSeeEurope                    // 2^5 = 32
	canSeeNorthAmerica              // 2^6 = 64
	canSeeSourthAmerica             // 2^7 = 128
)

func main() {
	// not allow to set a constant that require something run at runtime
	// const piConst float64 = math.Sin(1.57)
	const myConst int = 42
	fmt.Printf("%+v %T\n", myConst, myConst)
	fmt.Printf("%+v %T\n", a, a) // 0
	fmt.Printf("%+v %T\n", b, b) // 1
	fmt.Printf("%+v %T\n", c, c) // 2

	fileSize := 4000000000.
	fmt.Printf("%.2fGB\n", fileSize/GB) // 3.73GB

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles) // 100101
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is Headquarters? %v\n", isHeadquarters&roles == isHeadquarters)
}

/*
 * SUMMARY
 * 	Immutable, but can be shadowed
 *	Replaced by the compiler at compile time
 *		value must be calculable at compile time
 *	Named like variables
 *		PascalCase for exported constants
 *		camelCase for internal constants
 *	Typed constants work like immutable variables
 *		can interoperate only with same type
 *	Untyped constants work like literals
 *		can interoperate with similar types
 *	Enumerated constants
 *		special symbol iota allows related constants to be created easily
 *		iota starts at 0 in each const block and increments by one
 *		watch out of constant values that match zero values for variables
 *	Enumerated expressions
 *		operations that can be determined at compile time are allowed
 *			arithmetic
 *			bitwise operations
 *			bitshifting
 */
