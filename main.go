package main

import (
	"fmt"

	"github.com/bluele/gcache"
)

func main() {
	gc := gcache.New(5). // size odpowiada za ilość Setów które mogę mieć na daną chwilę gdy jedne mamy error przy wywołaniu już 2 działa na kluczach tylko wartością zmieniam wartość niezależnie ilość się nie zmienia cacha
				LRU(). // how it pop out the cache przyda się jak będę ją zwalniał zwalniamy najmniej używaną tutaj w tym wywołujemy get więc key jest używane więcej razy
				Build()
	gc.Set("key", "ok")
	gc.Set("key1", "no")
	gc.Set("key2", "ok")
	gc.Set("key3", "ok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	gc.Set("key4", "ok")
	gc.Set("key5", "ok")
	gc.Set("key6", "hallo")

	fmt.Println(gc.Keys(false))

	value1, err := gc.Get("key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value1)
	/*
		gc.Remove("key") // usuwanie manualnej
		value, err = gc.Get("key")
		if err != nil {
			panic(err)
		}
		fmt.Println("Get:", value)
	*/
}
