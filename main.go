package main

import (
	"fmt"
	"time"

	"github.com/bluele/gcache"
)

func main() {

	/*
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

		fmt.Println(gc.Keys(true)) // bool checkes if the time expired, a nie czy nie mieści się już w pamięci (size)

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
	*/

	var evictCounter, loaderCounter, purgeCounter int
	gc := gcache.New(20).
		LRU().
		LoaderExpireFunc(func(key interface{}) (interface{}, *time.Duration, error) {
			loaderCounter++
			expire := 1 * time.Second
			return "ok", &expire, nil
		}).
		EvictedFunc(func(key, value interface{}) {
			evictCounter++
			fmt.Println("evicted key:", key)
		}).
		PurgeVisitorFunc(func(key, value interface{}) {
			purgeCounter++
			fmt.Println("purged key:", key)
		}).
		Build()
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	time.Sleep(1 * time.Second)
	value, err = gc.Get("key")
	if err != nil {
		panic(err)
	}
	fmt.Println("Get:", value)
	gc.Purge()
	if loaderCounter != evictCounter+purgeCounter {
		panic("bad")
	}
}
