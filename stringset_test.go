package stringset

import (
	"fmt"
	"sort"
)

func Example_stringset_Delete() {
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

func Example_stringset_Add() {
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

func Example_stringset_Exists() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	element := "pippo"
	if ok := testSet.Exists(element); ok {
		fmt.Printf("%s exists", element)
	}
	// Output:
	// pippo exists
}

func Example_stringset_Strings() {
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

func Example_stringset_Strings_2() {
	testSet := NewStringSet()
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	//
}

func Example_stringset_Contains() {
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

func Example_stringset_Union() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := NewStringSet("pippo", "pluto", "minnie")

	u := testSet.Union(testSet2)

	slice := u.Strings()

	for _, element := range slice {
		fmt.Println(element)
	}
	// Unordered output:
	// minnie
	// paperino
	// pippo
	// pluto
}

func Example_stringset_Len() {
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

func Example_stringset_Pop() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")

	num := testSet.Len()
	for i := 0; i <= num; i++ { //testSet.Len() cannot be used in for loops
		element, _ := testSet.Pop()
		fmt.Println(element)
	}
	empty, ok := testSet.Pop()
	fmt.Println(empty, ok)
	// Unordered output:
	// paperino
	// pippo
	// pluto
	//
	//  false
}

func Example_stringset_Difference() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := NewStringSet("paperino", "pluto")

	diff := testSet.Difference(testSet2)

	fmt.Println(diff.Strings()[0])
	// Output:
	// pippo
}

func Example_stringset_Intersect() {
	testSet := NewStringSet("pippo", "pluto", "paperino", "pippo", "poldo", "minnie")
	testSet2 := NewStringSet("paperino", "pluto", "nonna papera")

	inersect := testSet.Intersect(testSet2)

	list := inersect.Strings()

	for _, element := range list {
		fmt.Println(element)
	}
	// Unordered output:
	// paperino
	// pluto
}

func Example_stringset_Intersect_2() {
	testSet := NewStringSet("paperino", "pluto", "nonna papera")
	testSet2 := NewStringSet("pippo", "pluto", "paperino", "pippo", "poldo", "minnie")

	inersect := testSet.Intersect(testSet2)

	list := inersect.Strings()

	for _, element := range list {
		fmt.Println(element)
	}
	// Unordered output:
	// paperino
	// pluto
}
