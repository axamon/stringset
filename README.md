# stringset

[![GoDoc](https://godoc.org/github.com/axamon/stringset?status.svg)](https://godoc.org/github.com/axamon/stringset)
[![Build Status](https://travis-ci.org/axamon/stringset.svg?branch=master)](https://travis-ci.org/axamon/stringset)
[![Go Report Card](https://goreportcard.com/badge/github.com/axamon/stringset)](https://goreportcard.com/report/github.com/axamon/stringset)
[![codecov](https://codecov.io/gh/axamon/stringset/branch/master/graph/badge.svg)](https://codecov.io/gh/axamon/stringset)


stringset creates sets for strings in golang that are concurrency safe



## Installation
```go get -u github.com/axamon/stringset```

## Usage
```go

package main

import (
    "fmt"
    "github.com/axamon/stringset"
)

func main() {
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
```


