# stringset

[![GoDoc](https://godoc.org/github.com/axamon/stringset?status.svg)](https://godoc.org/github.com/axamon/stringset)


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

    fmt.Println(testSet.Strings())

}
```


