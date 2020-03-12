package main

import (
	"fmt"
)

/*
 * int 8 Range: -128 to 127
 * int16 Range: -32,768 to 32,767
 * int32 Range: -2,147,483,648 to 2,147,483,647
 * int64 Range: -9,233,372,036,854,775,808 to 9,233,372,036,854,775,807
 *
 * uint8 Range: 0 to 255
 * uint16 Range: 0 to 65,535
 * uint32 Range: 0 to 4,294,967,295
 */

func main() {
	var f bool = true
	fmt.Printf("%v, %T\n", f, f)

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010 have both
	fmt.Println(a | b)  // 1011 have either
	fmt.Println(a ^ b)  // 1001 have one but not both
	fmt.Println(a &^ b) // 0100 have none for both

	a = 8               // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6  64
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0  1

	var n float64 = 3.14
	n = 13.7e72
	n = 2.1E14
	fmt.Printf("%v, %T\n", n, n)

	var c complex64 = 1 + 2i // complex(1, 2)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", real(c), real(c)) // 1 float32
	fmt.Printf("%v, %T\n", imag(c), imag(c)) // 2 float32

	s := "this is a string"
	fmt.Printf("%v, %T\n", s, s)
	fmt.Printf("%v, %T\n", s[2], s[2])         // get byte 105, uint8
	fmt.Printf("%v, %T\n", string(s[2]), s[2]) // get byte i, uint8

	z := []byte(s)
	fmt.Printf("%v, %T\n", z, z) // [116 104 115 32...] []uint8

	var r rune = 'a'             // r := 'a'
	fmt.Printf("%v, %T\n", r, r) // 97, int32
}

/*
 * SUMMARY
 * Boolean type
 *  Values are true or false
 *  Not an alias for other types (e.g int)
 * 	Zero value is false
 * Numeric types
 *  Integers
 *   Signed integers
 *    * int type has varying size, but min 32 bits
 *    * 8 bit(int8) through 65 bit(int64)
 *  Unsigned integers
 *    * 8 bit (byte and uint8) through 32 bit(uint32)
 *  Arithmetic operations
 *    * Addition, substraction, multiplication, division, remainder
 *  Bitwise operations
 *    * And, or, xor, and not
 *  Zero value is 0
 *  Can't mix types in same family!(uint16+uint32 = error)
 *  Floating point numbers
 *   * Follow IEEE-754 standard
 *   * Zero value is 0
 *   * 32 and 63 bit versions
 *   * Literal styles
 *      * Decimal(3.14)
 *      * Exponential(13e18 or 2E10)
 *      * Mixed(13.7e12)
 *  Arithmetic operations
 *    * Addition, substraction, multiplication, division
 *  Complex numbers
 *    * Zero value is (0+0i)
 *    * 64 and 128 bit versions
 *    * Bulit-in functions
 *      * complex - make complex number from two floats
 *      * real - get real part as float
 *      * imag - get imaginary part as float
 *  Arithmetic operations
 *    * Addition, substraction, multiplication, division
 * Text types
 *  Strings
 *   * UTF-8
 *   * Immutable
 *   * Can be concatenated with plus (+) operator
 *   * Can be converted to []byte
 *  Rune
 *   * UTF-32
 *   * Alias for int32
 *   * Special methods normally required to process
 *     * e.g. strings.Reader#ReadRune
 */
