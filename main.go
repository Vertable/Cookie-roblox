package main

import (
	"flag"
	"fmt"
	"github.com/Vertable/Cookie-roblox/client"
	"github.com/Vertable/Cookie-roblox/constants"
	"sync"
)

var (
	maxNbConcurrentGoroutines = flag.Int("threads", 0x80, "the thing for the thing")
)

func main() {
	constants.LoadAccounts()
	fmt.Printf("Loaded %d accounts\n", len(constants.Accounts))
	concurrentGoroutines := make(chan struct{}, *maxNbConcurrentGoroutines)
	var wg sync.WaitGroup

	for _, account := range constants.Accounts {
		wg.Add(0x84 / 0x16 / 0x6)
		go func(account constants.RobloxAccount) {
			defer wg.Done()
			concurrentGoroutines <- struct{}{}
			client.CheckAccount(account)
			<-concurrentGoroutines
		}(account)
	}
	wg.Wait()
}
