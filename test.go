package main

import (
	"fmt"
)

func main() {
	// intro
	fmt.Println("Hello, world")
	fmt.Println()

	// float
	var i  float32 = 42.5
	fmt.Printf("%v, %T\n", i, i)
	fmt.Println()

	// boolean
	n := 1 == 1
	m := 1 == 2
	fmt.Printf("%v, %T\n", n, n)
	fmt.Printf("%v, %T\n", m, m)

	var d bool
	fmt.Printf("%v, %T\n", d, d)
	fmt.Println()


	//unsigned integer
	var uintExample uint16 = 18
	fmt.Printf("%v, %T\n", uintExample, uintExample)
	fmt.Println()


	number1 := 8
	number2 := 7

	// AND, OR, bit shifting
	fmt.Println(number1 & number2)
	fmt.Println(number1 | number2)
	fmt.Println(number1 >> 3) // 2ˆ3 / 2ˆ3 = 2ˆ0 = 1
	fmt.Println(number1 << 3) // 2ˆˆ * 2ˆˆ = 2ˆ6 = 64

	// complex numbers
	var complexNumber complex128 = 436 + 2.4i
	fmt.Println(complexNumber)
	fmt.Printf("%v, %T\n", real(complexNumber), real(complexNumber))
	fmt.Printf("%v, %T\n", imag(complexNumber), imag(complexNumber))
	fmt.Println()


	// string (UTF-8)
	string1 := "pepyaka"
	string2 := "ololo"
	fmt.Printf("%v, %T\n", string1[4], string1[4]) // printing bytes since strings in Go are aliases for bytes
	fmt.Printf("%v, %T\n", string(string1[4]), string(string1[4])) // typacasting to a string
	fmt.Printf("%v, %T\n", string1 + string2, string1 + string2) // concatenating two strings

	string2UTF8 := []byte(string1)
	fmt.Printf("%v, %T\n", string2UTF8, string2UTF8) // string to ASCII/UTF-8 values (uint8)
	fmt.Println()

	// rune (UTF-32)
	runeExample := 'a'
	fmt.Printf("%v, %T\n", runeExample, runeExample)
	fmt.Println()

	/* constants: are immutable, but can be shadowed; value must be calculable at compile time; same naming rules like for variables;
	typed constants work like immutable vars, but can only interoperate with the same type; untyped constants work like literals, and can interoperate with similar types*/
	const myConst int  = 53
	fmt.Printf("%v, %T\n", myConst, myConst) //remember that inner constant declaration wins over package level declaration (outside of main(); package level constant shadows an inner one)
	var number3 = 38
	fmt.Printf("%v, %T\n", myConst + number3, myConst + number3) // addition with a constant

	// constant with iota counter
	const (
		a = iota // pattern of naming constants in the block; each subsequent constant is assigned an iota value, but iota is scoped to this block
		b
		c
	)
	fmt.Printf("%v\n", a) // iota increases its value with each subsequent constant in the block
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)

	const (
		a2 = iota
	)
	fmt.Printf("%v\n", a2) // proof that the previous iota (a) is scoped to the previous block
	fmt.Println()

	// checking if a value has been assigned to a constant yet
	const (
		errorConst = iota // int == 0
		firstConst // int == 1
		secondConst // int == 2
		thirdConst // int == 3
	)

	var constType int // not defining a value == default
	fmt.Printf("%v\n", constType == firstConst) // false because iota assumes default value at the first const — errorConst

	// bitshifting with constants
	const (
		_ = iota // ignoring the first vaule
		KB =  1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fileSize := 4000000000.
	fmt.Printf("%.2fGB", fileSize/GB)
	fmt.Println()
	fmt.Println()

	// bitshifting for role storage & checks
	const (
		isAdmin = 1 << iota
		isHeadquarters
		canSeeFinance

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica

	)
	var roles byte = isAdmin | canSeeFinance | canSeeEurope
	fmt.Printf("%b\n", roles) // showing that data is encoded into a byte
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin) // checking Admin role — true
	fmt.Printf("Is HQ? %v\n", isHeadquarters & roles == isHeadquarters) // checking isHeadquaters — false
	fmt.Println()

	// arrays: their elements are contiguous in memory and faster to access; arrays are values in Go, not references; when array is copied, it's not pointing at the same underlying data, but a different set of data
	grades := [3]int{34, 57, 68} // fixed-size array
	fmt.Printf("Grades: %v\n", grades)
	dynamicSizeGrades := [...]int{34, 57, 68, 75, 99, 88} // dynamic-size array
	fmt.Printf("Grades: %v\n", dynamicSizeGrades)
	var students [3]string // empty array
	fmt.Printf("Students: %v\n", students)
	students[0] = "Sofia" //dynamic value insertion
	fmt.Printf("Students: %v\n", students)
	students[1] = "Anastasia"
	fmt.Printf("Student #1: %v\n", students[1])
	fmt.Printf("Number of students: %v\n", len(students))
	fmt.Println()

	// identity matrix
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{34, 45, 57}
	identityMatrix[1] = [3]int{46, 68, 27}
	identityMatrix[2] = [3]int{457, 37, 235}
	fmt.Println( identityMatrix)

	// copying an array
	firstArray := [...]int{1, 2, 3}
	secondArray := firstArray
	secondArray[1] = 10
	fmt.Println(firstArray) // original values
	fmt.Println(secondArray) //different second value, which means secondArray is a literal copy of firstArray and not pointing to firstArray; this potentially slows the running down

	//pointing to an array
	thirdArray := [...]int{1, 2, 3}
	forthArray := &thirdArray // forthArray is pointing to thirdArray
	forthArray[1] = 10
	fmt.Println(thirdArray) // [1] value has changed, which proves forthArray is pointing to thirdArray
	fmt.Println(forthArray)
	fmt.Println()

	// slices: they are a reference type; slicing operations on slices and arrays; slices cannot be checked for equality;  a slice cannot be a key to a map
	firstSlice := []int{1, 2, 3}
	fmt.Println(firstSlice)
	fmt.Printf("Length of slice: %v\n", len(firstSlice))
	fmt.Printf("Capacity of slice: %v\n", cap(firstSlice))
	secondSlice := firstSlice
	secondSlice[2] = 3658
	fmt.Println(firstSlice) // [2] value changed, which proves secondSlice is pointing to firstSlice
	fmt.Println(secondSlice)
	fmt.Println()

	// slicing slices :-)
	thirdSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	forthSlice := thirdSlice[:] // all elements
	fifthSlice := thirdSlice[4:] // 4-th element on (not including 4-th)
	sixthSlice := thirdSlice[:7] // up to 7-th element (including 7-th)
	seventhSlice := thirdSlice[2:8] // from 2-nd up to (not including 2-nd) 8-th elements
	fmt.Println(forthSlice)
	fmt.Println(fifthSlice)
	fmt.Println(sixthSlice)
	fmt.Println(seventhSlice)
	fmt.Println()

	// making slices
	eightSlice := make([]int, 3, 100) /* create a 3 element slice capacity == 100 elements; make is handy to reduce memory consumption while dynamically appending values to a slice, as we can initialize a slice with a capacity that we plan for the future
	because if the capacity is exceeded, a slice is fully copied when appending elements; otherwise, we work with the initial slice via a pointer */
	fmt.Println(eightSlice) // print default slice values (initialized to 0)
	fmt.Printf("Length: %v\n", len(eightSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(eightSlice)) // print slice capacity
	fmt.Println()


	ninthSlice := []int{} // slices are dynamic; let's initialize a slice with 0 elements
	fmt.Println(ninthSlice) // print default slice values
	fmt.Printf("Length: %v\n", len(ninthSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	ninthSlice = append(ninthSlice, 1) // let's append an element, i.e. make a full copy of ninthSlice and add an elements to it
	fmt.Println(ninthSlice) // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice)) // print slice length that changed from 0 to 1
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	fmt.Println()

	// stack operations with a slice — append
	ninthSlice = append(ninthSlice, 2, 3, 4, 5, 6, 7) // we can append more than 1 element at a time, but of the slice type; i.e. we cannot add []int{2, 3, 4, 5, 6, 7} a slice of integers, but only integers
	fmt.Println(ninthSlice) // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	fmt.Println()

	// appending a slice to a slice (workaround)
	ninthSlice = append(ninthSlice, []int{8, 9, 10}...) // but we can use this workaround — Go is going to decompose the appended slice to individual elements
	fmt.Println(ninthSlice) // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	fmt.Println()

	//  stack operations with a slice — remove
	tenthSlice := ninthSlice[1:] // trim the first elements by shifting
	fmt.Println(tenthSlice) // print new slice
	fmt.Printf("Length: %v\n", len(tenthSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(tenthSlice)) // print slice capacity
	fmt.Println()

	eleventhSlice := ninthSlice[:len(ninthSlice)-1] // trim the last element
	fmt.Println(eleventhSlice) // print new slice
	fmt.Printf("Length: %v\n", len(eleventhSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(eleventhSlice)) // print slice capacity
	fmt.Println()

	fmt.Println("This is the initial slice before removing the 3rd element:", ninthSlice) // we change the initial array, i.e. twelfthSlice is pointing to ninthSlice
	twelfthSlice := append(ninthSlice[:2], ninthSlice[3:]...) // remove elements that are in other position — 3rd element in this example
	fmt.Println("This is the initial slice after removing the 3rd element:", ninthSlice) // we change the initial slice adding a new slice and the last value is duplicated; so remember not to have any other pointers to the same slice to avoid unexpected havoc
	fmt.Println("This is the new slice", twelfthSlice) // print new slice
	fmt.Printf("Length: %v\n", len(twelfthSlice)) // print slice length
	fmt.Printf("Capacity: %v\n", cap(twelfthSlice)) // print slice capacity
	fmt.Println()

	// maps: maps cannot be checked for equality; a slice cannot be a key to a map
	statePopulations := map[string]int{
		"CA": 39250017,
		"TX": 27862596,
		"FL": 20612439,
		"NY": 19745289,
	} // map with strings as keys and integers as values
	fmt.Println(statePopulations)
	fmt.Println()

	// firstMap := map[[]int]string{} as said a slice cannot be a key to a map

	firstMap := map[[3]int]string{} // but we can turn a slice into an array :-)
	fmt.Println(statePopulations, firstMap) // print two maps; firstMap is, of course, empty
	fmt.Println()

	// let's initialize another empty map for future entries, via make
	secondMap := make(map[string]int)
	fmt.Println(secondMap)
	fmt.Println()


	fmt.Println(statePopulations["NY"]) // let's print a value from statePopulations
	fmt.Println(statePopulations) // before adding GA
	statePopulations["GA"] = 10310371 // let's add a value to statePopulations
	fmt.Println(statePopulations["GA"]) // let's print GA
	fmt.Println(statePopulations) // after adding GA
	delete(statePopulations, "GA")// let's delete GA from statePopulations
	fmt.Println(statePopulations) // after deleting GA
	fmt.Println()

	// maps are addressed by reference, so changes affect the source
	thirdMap := statePopulations // thirdMap points to statePopulations
	fmt.Println(statePopulations) // statePopulations before deleting NY from thirdMap
	delete(thirdMap, "NY")
	fmt.Println(thirdMap) // thirdMap after deleting NY from it
	fmt.Println(statePopulations) // statePopulations after deleting NY from thirdMap
	fmt.Println()


	// structs: are value types, not reference types; we cannot publish field names, i.e. packages will only be able to see the struct, but not its field names
	type Colleague struct {
		number int
		name string
		colleagues []string
	}

	aColleague := Colleague {
				number: 1,
				name: "Sofia",
				colleagues: []string {
					"Dan",
					"Vlad",
					"Cyrill",
				},
	} /* if we use field names, they can be in whatever order in a struct and Go will map them by a name; e.g.
	aColleague := Colleague {
			colleagues: []string {
				"Dan",
				"Vlad",
				"Cyrill",
			},
			name: "Sofia",
			number: 1,
	and if we use positional syntax (no field names, only values), we'll quickly run into problems when new field are added to the struct */
	fmt.Println(aColleague) // print the whole struct
	fmt.Println(aColleague.name) // print an element from the struct
	fmt.Println(aColleague.colleagues) // print the "colleagues" slice
	fmt.Println(aColleague.colleagues[1]) // print only the second colleagues from the "colleagues" slice
	fmt.Println()


	anotherColleague := aColleague
	anotherColleague.name = "Peter" // as structs are value types, this change is going to affect only anotherColleague
	fmt.Println(anotherColleague)
	fmt.Println(aColleague) // aColleague remains unchanged

	oneMoreColleague := &aColleague
	oneMoreColleague.name = "Peter" // however, if we use a pointer...
	fmt.Println(oneMoreColleague)
	fmt.Println(aColleague) // ...aColleague changes its corresponding field value; well, it's pointer, what did you expect... :-)
	fmt.Println()

	// composition: Go doesn't have classic OOP inheritance, so a struct cannot inherit another struct, i.e. structs are independent; but a struct can have (characteristic of) another struct
	type Animal struct {
		Name string
		Origin string
	}

	type Cat struct {
		Animal
		speedKPH float32
		canMeow bool
		canDropThingFromSurfaces bool
		canAskForFood bool
	} // here we embed Animal into Cat
	gingerCat := Cat{}
	gingerCat.Name = "Tom"
	gingerCat.Origin = "US"
	gingerCat.canMeow = true
	gingerCat.canDropThingFromSurfaces = true
	gingerCat.canAskForFood = true
	fmt.Println(gingerCat) // proof that a cat has an animal (maybe its owner) :-)
	fmt.Println()

	// tags in structs: we can set tags for fields; we need to import Go reflection package for it ("reflect")
	/* type Mammal struct {
		Name string 'required max: "100"'
		Origin string
	} . */
}
