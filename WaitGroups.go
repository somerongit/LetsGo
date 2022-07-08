package main

import "sync"

func runWaitGroups() {

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		println("1st Routine...")
		wg.Done()
	}()

	go func() {
		println("2nd Routine...")
		wg.Done()
	}()

	wg.Wait()
	println("Next process...")
	// select {}
	// fatal error: all goroutines are asleep - deadlock!
}
