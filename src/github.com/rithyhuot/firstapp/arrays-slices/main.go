package main

import (
	"fmt"
)

func main() {
	grades := [3]int{97, 85, 93} // also can do [...]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades)

	var students [3]string
	fmt.Printf("Students: %v\n", students) // Students: []
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)                // Students: [Lisa]
	fmt.Printf("Number of Students: %v\n", len(students)) // Number of Students: 3

	a := [...]int{1, 2, 3}
	b := a // make a copy
	b[1] = 5
	fmt.Println(a) // [1,2,3]
	fmt.Println(b) // [1,5,3]

	c := &a // c points to a
	c[1] = 10
	fmt.Println(a) // [1,10,3]
	fmt.Println(c) // &[1,10,3]

	d := []int{1, 2, 3}                     // slice
	f := d                                  // f points to d
	fmt.Printf("%+v, %v\n", len(d), cap(d)) // 3, 3
	fmt.Printf("%+v, %v\n", len(f), cap(f)) // 3, 3

	aa := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	bb := aa[:]   // slice of all elements
	cc := aa[3:]  // slice from 4th element to end
	dd := aa[:6]  // slice first 6 elements
	ee := aa[3:6] // slice the 4th, 5th and 6th elements

	fmt.Println(aa) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(bb) // [1 2 3 4 5 6 7 8 9 10]
	fmt.Println(cc) // [4 5 6 7 8 9 10]
	fmt.Println(dd) // [1 2 3 4 5 6]
	fmt.Println(ee) // [4 5 6]

	aaa := make([]int, 3, 100)    // type, length, capacity of the slice
	fmt.Printf("%+v\n", aaa)      // [0, 0, 0]
	fmt.Printf("%+v\n", len(aaa)) // 3
	fmt.Printf("%+v\n", cap(aaa)) // 100

	bbb := []int{}                          // slice of int with no element
	fmt.Printf("%+v\n", bbb)                // []
	fmt.Printf("%+v\n", len(bbb))           // 0
	fmt.Printf("%+v\n", cap(bbb))           // 0
	bbb = append(bbb, 1)                    // create new array
	fmt.Printf("%+v\n", bbb)                // [1]
	fmt.Printf("%+v\n", len(bbb))           // 1
	fmt.Printf("%+v\n", cap(bbb))           // 1
	bbb = append(bbb, 2, 3, 4, 5)           // create new array
	fmt.Printf("%+v\n", bbb)                // [1 2 3 4 5]
	fmt.Printf("%+v\n", len(bbb))           // 1
	fmt.Printf("%+v\n", cap(bbb))           // 1
	bbb = append(bbb, []int{6, 7, 8, 9}...) // have to spread the array since slice can only work with the same type int != int[]
	fmt.Printf("%+v\n", bbb)                // [1]
	fmt.Printf("%+v\n", len(bbb))           // 1
	fmt.Printf("%+v\n", cap(bbb))           // 1

	ccc := []int{1, 2, 3, 4, 5}
	ddd := ccc[1:]                     // shift off the first element and create new slice
	fff := ccc[:len(ccc)-1]            // remove last element and create new slice
	ggg := append(ccc[:2], ccc[3:]...) // remove an element in the middle of the slice and create new slice
	fmt.Printf("%+v\n", ddd)
	fmt.Printf("%+v\n", fff)
	fmt.Printf("%+v\n", ggg)
}

/*
 * SUMMARY
 *	Arrays
 *		collection of items with same type
 *		fixed size
 *		declaration stypes
 *			a := [3]int{1,2,3}
 *			a := [...]int{1,2,3}
 *			var a [3]int
 *		access via zero-based index
 *			a := [3]int{1,3,5} //a[1] == 3
 *		len function returns size of array
 *		copies refer to different underlying data
 *	Slices
 *		backed by array
 *		creation styles
 *			slice existing array or slice
 *			literal style
 *			via make function
 *				a := make([]int, 10) // create slice with capacity and length == 10
 *				a := make([]int, 10, 100) // slice with length == 10 and capacity == 100
 *		len function returns length of slice
 *		cap function returns to length of underlying array
 *		append function to add elements to slice
 *			may cause expensive copy operation if underlying array is too small
 *		copies refer to same underlying array
 */
