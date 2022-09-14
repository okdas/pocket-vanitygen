package main

import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/paulbellamy/ratecounter"
	"github.com/pokt-network/pocket-core/crypto"
)

func generateKeyAddressStrings() (string, string) {
	privKey := crypto.PrivateKey(crypto.Ed25519PrivateKey{}).GenPrivateKey()
	return privKey.RawString(), hex.EncodeToString(privKey.PubKey().Address().Bytes())
}

func main() {
	patternsPtr := flag.String("patterns", "", "look for patterns (comma separated). Pattern can have beginning-end (dash separated, dash always must be provided e.g. 420- or -69)")
	cpusPtr := flag.Int("cpus", runtime.NumCPU(), "how many cpus to use")
	flag.Parse()

	patternsSeparated := strings.Split(*patternsPtr, ",")
	patterns := make([][]string, len(patternsSeparated))
	for ind, pattern := range patternsSeparated {
		patterns[ind] = strings.Split(pattern, "-")
	}

	if len(patterns) == 0 {
		fmt.Println("No patterns provided")
		os.Exit(1)
	}

	fmt.Println("Looking for: ", patterns)
	fmt.Println("running on cpus: ", *cpusPtr)

	var wg sync.WaitGroup
	counter := ratecounter.NewRateCounter(1 * time.Second)
	go func() {
		time.Sleep(10 * time.Second)
		for {
			fmt.Println("Rate is", counter.Rate(), "per second")
			time.Sleep(1 * time.Hour)
		}
	}()

	for i := 0; i < *cpusPtr; i++ {
		wg.Add(1)
		go func() {
			for {
				for i := 0; i < 1000; i++ {
					counter.Incr(1000)
					pk, addr := generateKeyAddressStrings()
					for _, pattern := range patterns {
						if strings.HasPrefix(addr, pattern[0]) && strings.HasSuffix(addr, pattern[1]) {
							fmt.Println("Found", addr, pk)
						}
					}
				}

			}
		}()
	}
	wg.Wait()
}
