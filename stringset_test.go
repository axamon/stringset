package stringset

import (
	"fmt"
)

func ExampleDelete() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	testSet.Delete("pluto")
	testSet.Delete("nonna papera")
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	// pippo
	// paperino

}

func Example_Add() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	testSet.Add("pluto")
	testSet.Add("nonna papera")
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	// pippo
	// pluto
	// paperino
	// nonna papera

}

func Example_Exists() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	element := "pippo"
	if ok := testSet.Exists(element); ok {
		fmt.Printf("%s exists", element)
	}
	// Output:
	// pippo exists
}

func ExampleStrings() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	testSet = NewStringSet()
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	// pippo
	// pluto
	// paperino
	//
}
