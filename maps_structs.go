package main

import (
	"fmt"
	"reflect"
)

// map 
func main() {
	statePopulations := make(map[string]int)
	statePopulations = map[string]int{
		"California": 39250017,
		"Texas": 27862596,
		"Florida": 20612439,
		"New York": 19745289,
		"Pennsylvania": 12802503,
		"Illınois": 12801539,
		"Ohio": 11614373,
	}
	statePopulations["Georgia"] = 1031524 // Add element
	delete(statePopulations, "Georgia") // Delete element
	fmt.Println(statePopulations["Ohio"]) // Pull up one information from map

	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)
}

// Struct

type Doctor struct {
	number int
	actorName string
	companions []string
}

func example() {
	aDoctor := Doctor{
		number: 3,
		actorName: "Katy Perry",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
		},
	}
	fmt.Println(aDoctor)
}

// Embedding

type Animal struct {
	Name string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKMH float32
	CanFly bool
}

func embed() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKMH = 48
	b.CanFly = true
	fmt.Println(b)

}

// Reflect Package

func ref() {
	t := reflect.TypeOf(Animal{})
	field, _ :=t.FieldByName("Name")
	fmt.Println(field.Tag)	
}

