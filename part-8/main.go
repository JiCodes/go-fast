package main

import (
	"fmt"
	"time"
	"sync"
	
)

// goroutine
// let us launch multiple functions and have them execute concurrently

// cocurrency is not parallelism

var wg = sync.WaitGroup{} // wait group
var mutex = sync.Mutex{} // mutual exclusion
var m = sync.RWMutex{} // read write mutex

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{} // slice to store the results

func main() {
	t0 := time.Now()
	for i:= range dbData {
		wg.Add(1) // add 1 to wait group
		go dbCall(i) // go routine
	}
	wg.Wait() // wait for all the go routines to finish
	fmt.Printf("Total execution time: %v\n", time.Since(t0))
	fmt.Printf("The results are: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])

	// mutex.Lock() // lock the mutex
	// results = append(results, dbData[i])
	// mutex.Unlock()

	save(dbData[i])
	log()

	wg.Done()
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	m.RLock()
	fmt.Printf("The current results are: %v\n", results)
	m.RUnlock()
}