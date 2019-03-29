package stringset_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/axamon/stringset"
)

var tt = stringset.NewStringSet()

func BenchmarkT(b *testing.B) {
	var t = stringset.NewStringSet()
	for n := 0; n < b.N; n++ {
		go t.Add("pippo")
		go t.Add("pluto")
		go t.Len()
		go t.Exists("pippo")
		go t.Delete("pippo")
		go t.Pop()
	}
}

func BenchmarkAdd(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go tt.Add("pippo")
	}
}

func BenchmarkDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		go tt.Delete("pippo")
	}
}

func BenchmarkIntersect(b *testing.B) {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo", "poldo", "minnie", "paperinik", "paperoga")
	testSet2 := stringset.NewStringSet("paperino", "pluto", "nonna papera")

	for n := 0; n < b.N; n++ {

		go testSet.Intersect(testSet2)

	}
}

func ExampleStringSet_Delete() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")

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

func ExampleStringSet_Add() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")

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

func ExampleStringSet_Exists() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")

	element := "pippo"
	if ok := testSet.Exists(element); ok {
		fmt.Printf("%s exists", element)
	}
	// Output:
	// pippo exists
}

func ExampleStringSet_Strings() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")
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

func ExampleStringSet_Strings_second() {
	testSet := stringset.NewStringSet()
	for _, element := range testSet.Strings() {
		fmt.Println(element)
	}
	// Output:
	//
}

func ExampleStringSet_Contains() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := stringset.NewStringSet("pippo", "pluto")

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

func ExampleStringSet_Union() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := stringset.NewStringSet("pippo", "pluto", "minnie")

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

func ExampleStringSet_Len() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := stringset.NewStringSet("pippo", "pluto")
	testSet3 := stringset.NewStringSet()

	fmt.Println(testSet.Len())
	fmt.Println(testSet2.Len())
	fmt.Println(testSet3.Len())
	// Output:
	// 3
	// 2
	// 0
}

func ExampleStringSet_Pop() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")

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

func ExampleStringSet_Difference() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo")
	testSet2 := stringset.NewStringSet("paperino", "pluto")

	diff := testSet.Difference(testSet2)

	fmt.Println(diff.Strings()[0])
	// Output:
	// pippo
}

func ExampleStringSet_Intersect() {
	testSet := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo", "poldo", "minnie")
	testSet2 := stringset.NewStringSet("paperino", "pluto", "nonna papera")

	inersect := testSet.Intersect(testSet2)

	list := inersect.Strings()

	for _, element := range list {
		fmt.Println(element)
	}
	// Unordered output:
	// paperino
	// pluto
}

func ExampleStringSet_Intersect_second() {
	testSet := stringset.NewStringSet("paperino", "pluto", "nonna papera")
	testSet2 := stringset.NewStringSet("pippo", "pluto", "paperino", "pippo", "poldo", "minnie")

	inersect := testSet.Intersect(testSet2)

	list := inersect.Strings()

	for _, element := range list {
		fmt.Println(element)
	}
	// Unordered output:
	// paperino
	// pluto
}
