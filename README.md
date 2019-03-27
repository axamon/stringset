# stringset

[![GoDoc](https://godoc.org/github.com/axamon/stringset?status.svg)](https://godoc.org/github.com/axamon/stringset)
[![Build Status](https://travis-ci.org/axamon/stringset.svg?branch=master)](https://travis-ci.org/axamon/stringset)


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
    testSet := stringset.NewStringSet("pippo", "pluto", "paperino","pippo")

    for _, element := range testSet.Strings() {
		fmt.Println(element)
	}

}
    // Output:
	// pippo
	// pluto
	// paperino
```


