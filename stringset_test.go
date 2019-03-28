package stringset

import (
	"fmt"
	"sort"
)

func ExampleDelete() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	testSet.Delete("pluto")
	slice := testSet.Strings()
	sort.Strings(slice)
	for _, element := range slice {
		fmt.Println(element)
	}
	// Output:
	// paperino
	// pippo
}

func Example_Add() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	testSet.Add("pluto")
	testSet.Add("nonna papera")
	slice := testSet.Strings()
	sort.Strings(slice)
	for _, element := range slice {
		fmt.Println(element)
	}
	// Output:
	// nonna papera
	// paperino
	// pippo
	// pluto
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
	slice := testSet.Strings()
	sort.Strings(slice)
	for _, element := range slice {
		fmt.Println(element)
	}
	// Output:
	// paperino
	// pippo
	// pluto
}

func ExampleStrings_2() {
	testSet := NewStringSet()
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	//
}

func Example_Contains() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := NewStringSet("pippo", "pluto")

	if ok := testSet.Contains(testSet2); ok {
		fmt.Println("Yes")
	}
	if ok := testSet2.Contains(testSet); !ok {
		fmt.Println("No")
	}
	// Output:
	// Yes
	// No
}

func Example_Union() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := NewStringSet("pippo", "pluto", "minnie")

	if ok := testSet.Union(testSet2); ok {
		fmt.Println("Done")
	}
	slice := testSet.Strings()
	sort.Strings(slice)
	for _, element := range slice {
		fmt.Println(element)
	}
	// Output:
	// Done
	// minnie
	// paperino
	// pippo
	// pluto
}

func Example_Len() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := NewStringSet("pippo", "pluto")
	testSet3 := NewStringSet()

	fmt.Println(testSet.Len())
	fmt.Println(testSet2.Len())
	fmt.Println(testSet3.Len())
	// Output:
	// 3
	// 2
	// 0
}
