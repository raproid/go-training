package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"reflect"
)

func main() {
	// intro
	fmt.Println("Hello, world")
	fmt.Println()

	// float
	var i float32 = 42.5
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
	fmt.Printf("%v, %T\n", string1[4], string1[4])                 // printing bytes since strings in Go are aliases for bytes
	fmt.Printf("%v, %T\n", string(string1[4]), string(string1[4])) // typacasting to a string
	fmt.Printf("%v, %T\n", string1+string2, string1+string2)       // concatenating two strings

	string2UTF8 := []byte(string1)
	fmt.Printf("%v, %T\n", string2UTF8, string2UTF8) // string to ASCII/UTF-8 values (uint8)
	fmt.Println()

	// rune (UTF-32)
	runeExample := 'a'
	fmt.Printf("%v, %T\n", runeExample, runeExample)
	fmt.Println()

	/* constants: are immutable, but can be shadowed; value must be calculable at compile time; same naming rules like for variables;
	typed constants work like immutable vars, but can only interoperate with the same type; untyped constants work like literals, and can interoperate with similar types */
	const myConst int = 53
	fmt.Printf("%v, %T\n", myConst, myConst) //remember that inner constant declaration wins over package level declaration (outside of main(); package level constant shadows an inner one)
	var number3 = 38
	fmt.Printf("%v, %T\n", myConst+number3, myConst+number3) // addition with a constant

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
		errorConst  = iota // int == 0
		firstConst         // int == 1
		secondConst        // int == 2
		thirdConst         // int == 3
	)

	var constType int                           // not defining a value == default
	fmt.Printf("%v\n", constType == firstConst) // false because iota assumes default value at the first const — errorConst

	// bitshifting with constants
	const (
		_  = iota // ignoring the first vaule
		KB = 1 << (10 * iota)
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
	fmt.Printf("%b\n", roles)                                         // showing that data is encoded into a byte
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)            // checking Admin role — true
	fmt.Printf("Is HQ? %v\n", isHeadquarters&roles == isHeadquarters) // checking isHeadquaters — false
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
	fmt.Println(identityMatrix)

	// copying an array
	firstArray := [...]int{1, 2, 3}
	secondArray := firstArray
	secondArray[1] = 10
	fmt.Println(firstArray)  // original values
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
	forthSlice := thirdSlice[:]     // all elements
	fifthSlice := thirdSlice[4:]    // 4-th element on (not including 4-th)
	sixthSlice := thirdSlice[:7]    // up to 7-th element (including 7-th)
	seventhSlice := thirdSlice[2:8] // from 2-nd up to (not including 2-nd) 8-th elements
	fmt.Println(forthSlice)
	fmt.Println(fifthSlice)
	fmt.Println(sixthSlice)
	fmt.Println(seventhSlice)
	fmt.Println()

	// making slices
	eightSlice := make([]int, 3, 100) /* create a 3 element slice capacity == 100 elements; make is handy to reduce memory consumption while dynamically appending values to a slice, as we can initialize a slice with a capacity that we plan for the future
	because if the capacity is exceeded, a slice is fully copied when appending elements; otherwise, we work with the initial slice via a pointer */
	fmt.Println(eightSlice)                       // print default slice values (initialized to 0)
	fmt.Printf("Length: %v\n", len(eightSlice))   // print slice length
	fmt.Printf("Capacity: %v\n", cap(eightSlice)) // print slice capacity
	fmt.Println()

	ninthSlice := []int{}                         // slices are dynamic; let's initialize a slice with 0 elements
	fmt.Println(ninthSlice)                       // print default slice values
	fmt.Printf("Length: %v\n", len(ninthSlice))   // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	ninthSlice = append(ninthSlice, 1)            // let's append an element, i.e. make a full copy of ninthSlice and add an elements to it
	fmt.Println(ninthSlice)                       // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice))   // print slice length that changed from 0 to 1
	fmt.Printf("Capacity: %v\n", cap(ninthSlice)) // print slice capacity
	fmt.Println()

	// stack operations with a slice — append
	ninthSlice = append(ninthSlice, 2, 3, 4, 5, 6, 7) // we can append more than 1 element at a time, but of the slice type; i.e. we cannot add []int{2, 3, 4, 5, 6, 7} a slice of integers, but only integers
	fmt.Println(ninthSlice)                           // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice))       // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice))     // print slice capacity
	fmt.Println()

	// appending a slice to a slice (workaround)
	ninthSlice = append(ninthSlice, []int{8, 9, 10}...) // but we can use this workaround — Go is going to decompose the appended slice to individual elements
	fmt.Println(ninthSlice)                             // print new slice values
	fmt.Printf("Length: %v\n", len(ninthSlice))         // print slice length
	fmt.Printf("Capacity: %v\n", cap(ninthSlice))       // print slice capacity
	fmt.Println()

	//  stack operations with a slice — remove
	tenthSlice := ninthSlice[1:]                  // trim the first elements by shifting
	fmt.Println(tenthSlice)                       // print new slice
	fmt.Printf("Length: %v\n", len(tenthSlice))   // print slice length
	fmt.Printf("Capacity: %v\n", cap(tenthSlice)) // print slice capacity
	fmt.Println()

	eleventhSlice := ninthSlice[:len(ninthSlice)-1]  // trim the last element
	fmt.Println(eleventhSlice)                       // print new slice
	fmt.Printf("Length: %v\n", len(eleventhSlice))   // print slice length
	fmt.Printf("Capacity: %v\n", cap(eleventhSlice)) // print slice capacity
	fmt.Println()

	fmt.Println("This is the initial slice before removing the 3rd element:", ninthSlice) // we change the initial array, i.e. twelfthSlice is pointing to ninthSlice
	twelfthSlice := append(ninthSlice[:2], ninthSlice[3:]...)                             // remove elements that are in other position — 3rd element in this example
	fmt.Println("This is the initial slice after removing the 3rd element:", ninthSlice)  // we change the initial slice adding a new slice and the last value is duplicated; so remember not to have any other pointers to the same slice to avoid unexpected havoc
	fmt.Println("This is the new slice", twelfthSlice)                                    // print new slice
	fmt.Printf("Length: %v\n", len(twelfthSlice))                                         // print slice length
	fmt.Printf("Capacity: %v\n", cap(twelfthSlice))                                       // print slice capacity
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

	firstMap := map[[3]int]string{}         // but we can turn a slice into an array :-)
	fmt.Println(statePopulations, firstMap) // print two maps; firstMap is, of course, empty
	fmt.Println()

	// let's initialize another empty map for future entries, via make
	secondMap := make(map[string]int)
	fmt.Println(secondMap)
	fmt.Println()

	fmt.Println(statePopulations["NY"]) // let's print a value from statePopulations
	fmt.Println(statePopulations)       // before adding GA
	statePopulations["GA"] = 10310371   // let's add a value to statePopulations
	fmt.Println(statePopulations["GA"]) // let's print GA
	fmt.Println(statePopulations)       // after adding GA
	delete(statePopulations, "GA")      // let's delete GA from statePopulations
	fmt.Println(statePopulations)       // after deleting GA
	fmt.Println()

	// maps are addressed by reference, so changes affect the source
	thirdMap := statePopulations  // thirdMap points to statePopulations
	fmt.Println(statePopulations) // statePopulations before deleting NY from thirdMap
	delete(thirdMap, "NY")
	fmt.Println(thirdMap)         // thirdMap after deleting NY from it
	fmt.Println(statePopulations) // statePopulations after deleting NY from thirdMap
	fmt.Println()

	// structs: are value types, not reference types;
	type Colleague struct {
		number     int
		name       string
		colleagues []string
	}

	aColleague := Colleague{
		number: 1,
		name:   "Sofia",
		colleagues: []string{
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
	fmt.Println(aColleague)               // print the whole struct
	fmt.Println(aColleague.name)          // print an element from the struct
	fmt.Println(aColleague.colleagues)    // print the "colleagues" slice
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
		Name   string
		Origin string
	}

	type Cat struct {
		Animal
		speedKPH                 float32
		canMeow                  bool
		canDropThingFromSurfaces bool
		canAskForFood            bool
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
	type Mammal struct {
		Name   string `required max: "100"`
		Origin string
	}
	tagExample := reflect.TypeOf(Mammal{})
	field, _ := tagExample.FieldByName("Name")
	fmt.Println(field.Tag) //print the tag
	fmt.Println()

	// if statements
	if pop, ok := statePopulations["FL"]; ok {
		fmt.Println(pop)
	} // pop is only defined and exists within the scope of the if statement
	fmt.Println()

	numbertoguess := 57
	guess := 39
	if guess < numbertoguess {
		fmt.Println("Too low")
	}
	if guess > numbertoguess {
		fmt.Println("Too high")
	}
	if guess == numbertoguess {
		fmt.Println("Spot on")
	}
	fmt.Println(numbertoguess <= guess, numbertoguess >= guess, numbertoguess != guess) // also checking smaller or equal to, greater or equal to, and not equal to
	fmt.Println()

	// logical tests for checks
	numbertoguess1 := 12
	guess1 := 45
	if guess1 < 1 || guess1 > 100 {
		fmt.Println("Your guess must be between 1 and 100.")
	} // an OR check for out-of-range guess values
	if guess1 >= 1 && guess1 <= 100 {
		if guess1 < numbertoguess1 {
			fmt.Println("Too low")
		}
		if guess1 > numbertoguess1 {
			fmt.Println("Too high")
		}
		if guess1 == numbertoguess1 {
			fmt.Println("Spot on")
		}
		fmt.Println(numbertoguess1 <= guess1, numbertoguess1 == guess1, numbertoguess1 >= guess1, numbertoguess1 != guess1)
	}
	fmt.Println()

	// logical tests where only one part runs. Either the OR or second part is executed
	numbertoguess2 := 24
	guess2 := 1
	if guess2 < 1 || guess2 > 100 {
		fmt.Println("Your guess must be between 1 and 100.")
	} else {
		if guess2 < numbertoguess2 {
			fmt.Println("Too low")
		}
		if guess2 > numbertoguess2 {
			fmt.Println("Too high")
		}
		if guess2 == numbertoguess2 {
			fmt.Println("Spot on")
		}
		fmt.Println(numbertoguess2 <= guess2, numbertoguess2 == guess2, numbertoguess2 >= guess2, numbertoguess2 != guess2)
	}
	fmt.Println()

	// comparing numbers here, we see these are the same numbers
	myNumber := 0.1
	if myNumber == math.Pow(math.Sqrt(myNumber), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("There are different")
	}
	fmt.Println()

	// here, however, we see these are different numbers since a floating point number is an approximation of decimal value, not an exact representation
	myNumber1 := 0.123
	if myNumber1 == math.Pow(math.Sqrt(myNumber), 2) {
		fmt.Println("These are the same")
	} else {
		fmt.Println("There are different")
	}
	fmt.Println()

	// switch statements

	// the value of a case is compared with the tag (part after the "switch" keyword) and the case is gonna execute if the value matches
	switch 2 {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Tswo")
	default:
		fmt.Println("Neither one or two")
	}
	fmt.Println()

	// Go allows for multiple tests in a single case. Naturally, overlapping cases when using multiple tests in a single case are not allowed in Go. So, an syntax error pops up in a situation like "case 1, 4, 9"  and "case 2, 6, 9".
	switch 3 {
	case 1, 4, 9:
		fmt.Println("One, four or nine")
	case 2, 6, 10:
		fmt.Println("Two, six or ten")
	default:
		fmt.Println("Another number")
	}
	fmt.Println()

	// Go allows for initializers. For example, "i=3+5" initializes the value of the tag "i" that follows it.
	switch i := 3 + 5; i {
	case 1, 4, 9:
		fmt.Println("One, four or nine")
	case 2, 6, 10:
		fmt.Println("Two, six or ten")
	default:
		fmt.Println("Another number")
	}
	fmt.Println()

	/* There is also a tagless syntax for switch cases. A variable declared before a switch statement.
	In a tagless syntax, cases are allowed to overlap (10 is lte 10 and also lte 20). If they do, the first case that evaluate to true, is gonna execute.
	The delimiter for the statemets in the case is the "case" keywords, "default" keyword or the closing brace, i.e. there can be multiple operations within a single case.
	The "break" keyword at the end of a case is implied and doesn't have to be explicitly stated. */
	s := 10
	switch {
	case s <= 10:
		fmt.Println("Less than or equal to ten")
	case s <= 20:
		fmt.Println("More than or equal to twenty")
	default:
		fmt.Println("Greater than twenty")
	}
	fmt.Println()

	// for falling through, Go offers the "fallthrough" keyword. So, both the first and second case will execute in the example below. An important thing is the keyword is logicless, so the second case executes even if it doesn't fit.
	g := 10
	switch {
	case g <= 10:
		fmt.Println("Less than or equal to ten")
		fallthrough
	case g <= 20:
		fmt.Println("More than or equal to twenty")
	default:
		fmt.Println("Greater than twenty")
	}
	fmt.Println()

	/* a good type switch example. j is a type interface that can take any type of data.
	The "break" keyword may be used explicitly to break out earlier than a case ends. E.g. a break can be wrapped in a logical test to determine a validation error in some incoming data, in which case this data should not be saved to the db.
	*/
	var j interface{} = 1
	switch j.(type) {
	case int:
		fmt.Println("j is an integer")
		break
		fmt.Println("This prints too")
	case float64:
		fmt.Println("j is a float")
	case string:
		fmt.Println("j is a string")
	default:
		fmt.Println("j is another type ")
	}
	fmt.Println()

	// looping

	//simple loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()

	//Go doesn't allows separating multiple statements with comma, but allows initializing multiple values at the same time
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
	fmt.Println()

	// looping even and odd numbers
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i%2 == 0 {
			i /= 2
		} else {
			i = 2*1 + 1
		}
		break
	}
	fmt.Println()

	//applying a custom tag inside the inner loop to break out the outer loop
Loop:
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Println(i * j)
			if i*j <= 4 {
				break Loop
			}
		}
	}
	fmt.Println()

	//using a loop for iterating over a map

	for k, v := range statePopulations {
		fmt.Println(k, v)
	}
	fmt.Println()

	//using a loop for iterating over a map printing only values and omitting keys

	for _, v := range statePopulations {
		fmt.Println(v)
	}
	fmt.Println()

	//using a loop for printing out a "Hello, Go!" with values as integers
	hello := "Hello, Go!"
	for k, v := range hello {
		fmt.Println(k, v)
	}
	fmt.Println()

	//using a loop for printing out a "Hello, Go!" and casting the values to chars
	for k, v := range hello {
		fmt.Println(k, string(v))
	}
	fmt.Println()

	//defer, panic, and recovery

	//defer executes any function passed into it after the (main() in this example) function finishes its final statement but before it returns
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	fmt.Println()

	/* deferred functions execute in the LIFO order, i.e. end-middle-start in this example...and the deferred "middle" from the previous example goes...last, just before the main() returns :D
	Deferred functions are often used to close out resources and the LIFO order is applied since one resource may depend on another one.
	*/
	defer fmt.Println("start")
	defer fmt.Println("middle")
	defer fmt.Println("end")
	fmt.Println()

	/* good deferring case is a program where we need to run some more logic after the request has been made and before the resource closes.
	We may actually forget to close the resource and the deferring it a neat solution in this case.
	Another good idea is to put the deferred resource closing right after the resource opening but not before checking for possible errors (for resource opening)
	*/
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("%s", string(robots))
	fmt.Println()

	/* deferred function may take the argument at the time the defer is called, not at the time the called function is executed. Hence, "start" is printed.
	..and the deferred "middle" from some preivous example goes...last, just before the main() returns :D */
	ac := "start"
	defer fmt.Println(ac)
	ac = "end"
	fmt.Println()

	// panic in Go is really a substitute for "exception" :D

	// dividing by zero calls a panic AAAAAAAAAAAAAAA
	// commented it not to...panic :D
	/* h, l := 1, 0
	answer := h / l
	fmt.Println(answer) */

	// a good panic example using a simple web handler. Well, couldn't make it work, so I commented it out.
	/* web_handler()
	fmt.Println()
	*/

	/* panic happens after the deferred statement is executed, so first deferring, then handling any panic, and only then handling the returned value.
	So, the deferred statements that we may use to close resources are going to succeed before the program panics.
	Again, commented the code not to get a panic.
	*/
	/* fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something bad happened")
	fmt.Println("end")
	fmt.Println()
	*/

}
