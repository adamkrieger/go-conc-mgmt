package main

import (
	"fmt"
)

type idNumber = string
type name = string
type idMap = map[idNumber]name

var (
	employees = idMap{
		"AK123": "Adam Krieger",
		"FB458": "Foo Bar",
	}
)

func main() {
	found, ref := identifyYourself("AK123")

	if found {
		fmt.Println(ref)
	}
}

func identifyYourself(id idNumber) (found bool, foundName name) {
	valueAtKeyThatMightBeNil := employees[id]

	valueAtKeyThatMightBeNil, trueIfWasFound := employees[id]

	return trueIfWasFound, valueAtKeyThatMightBeNil
}
