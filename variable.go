package main

import (
	"fmt"
	"strconv"
)

// Package Level Declaration
var i int = 42

// Package Level Block Declaration 
var (
	actorName string = "Elisabeth Sladen"
	companion string = "Sarah Jane Smith"
	doctorNumber int = 3
	season int = 11
)

// Different Types Of Declaration
func main() {
	var i int 
	i = 42
//	var j int = 27
//	k := 99
	fmt.Println(i)
}

// Shadowing
var a int = 27

func shadowing() {
	fmt.Println(a)
	var a int = 42
	fmt.Println(a)
}

// Variable Convertion

func convert() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i)
	fmt.Printf("%v, %T\n", j, j)
}

// String Convertion

func stringConvertion() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j string
	j = strconv.Itoa(i)
	fmt.Printf("%v, %T\n", j, j)
}