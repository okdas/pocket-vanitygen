// once all possible numbers are generated (can be checked with `ls found | wc -l`),
// run `cat found/*` to get all the found addresses

package main

import (
	"bufio"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"regexp"
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

type found struct {
	address string
	pk      string
	number  string
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
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

	founds := make(chan found, 100)

	go func() {
		for newItem := range founds {
			// fmt.Println("Found", newItem.number, newItem.address, newItem.pk)
			writeLines([]string{"\"" + newItem.number + "\": " + newItem.pk + " # " + newItem.address}, "found/"+newItem.number)
		}
	}()

	r := regexp.MustCompile("^(\\d{3})04")
	for i := 0; i < *cpusPtr; i++ {
		wg.Add(1)
		go func() {
			for {
				for i := 0; i < 1000; i++ {
					counter.Incr(1000)
					pk, addr := generateKeyAddressStrings()
					if r.MatchString(addr) {
						number := r.FindStringSubmatch(addr)[1]
						founds <- found{addr, pk, number}
					}
				}

			}
		}()
	}
	wg.Wait()
}
