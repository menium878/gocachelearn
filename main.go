package main

import (
	"fmt"

	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(5). // size odpowiada za ilość Setów które mogę mieć na daną chwilę gdy jedne mamy error przy wywołaniu już 2
				LRU(). // how it pop out the cache przyda się jak będę ją zwalniał
				Build()
	gc.Set("key", "ok")
	gc.Set("key1", "no")
	gc.Set("key", "ok")
	gc.Set("key", "ok")
	gc.Set("key", "ok")
	gc.Set("key", "ok")
	gc.Set("key", "hallo")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	value1, err := gc.Get("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value1)
	gc.Remove("key") // usuwanie manualnej
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
}
