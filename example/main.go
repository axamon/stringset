package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/axamon/stringset"
)

func main() {

	start := time.Now()

	tt1 := stringset.NewStringSet("pippo", "pluto", "paperino")

	var wg sync.WaitGroup

	for i := 0; i <= 100000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tt1.Add("topolino")
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			tt1.Delete("topolino")
		}()
	}
	wg.Wait()

	fmt.Println(tt1)

	elapsedtime := time.Since(start)
	fmt.Printf("elapsed time %3s\n", elapsedtime)

}
