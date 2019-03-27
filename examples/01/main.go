package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Pallinder/go-randomdata"

	"github.com/axamon/stringset"
)

func main() {

	start := time.Now()

	tt1 := stringset.NewStringSet("pippo", "pluto", "paperino")

	Heavyload(tt1, 10000)
	HeavyAdd(tt1, 10000)

	fmt.Println(tt1)

	elapsedtime := time.Since(start)
	fmt.Printf("elapsed time %3s\n", elapsedtime)

}

func HeavyAdd(tt1 *stringset.StringSet, num int) {
	var wg sync.WaitGroup

	for i := 0; i <= num; i++ {
		value := randomdata.FirstName(1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			tt1.Add(value)
		}()
	}
	wg.Wait()
	return
}

func Heavyload(tt1 *stringset.StringSet, num int) {
	var wg sync.WaitGroup

	for i := 0; i <= num; i++ {
		value := randomdata.FirstName(1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			tt1.Add(value)
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			tt1.Delete(value)
		}()
	}
	wg.Wait()
	return
}
