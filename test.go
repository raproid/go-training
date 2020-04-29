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

	// constants
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

}