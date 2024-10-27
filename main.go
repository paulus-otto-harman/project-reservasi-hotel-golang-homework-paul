package main

import (
	"homework/service"
	"homework/util"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go service.AppContainer(&wg, util.SessionTimeout)
	wg.Wait()

}
