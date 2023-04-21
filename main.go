package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ScarletTanager/redis-stream-balancer/rsb"
)

var (
	keyLen      int
	maxKeyCount int
	alphaChars  []rune = []rune("abcdefghijklmnopqrstuvwxyz")
)

func init() {
	flag.IntVar(&keyLen, "keyLength", 3, "Number of characters per key")
	flag.IntVar(&maxKeyCount, "maxKeyCount", 16384, "Maximum number of keys to compute")
}

func main() {
	flag.Parse()

	slots := make([][]string, 16384)
	for sp, _ := range slots {
		slots[sp] = make([]string, 0)
	}

	for _, p1 := range alphaChars {
		for _, p2 := range alphaChars {
			for _, p3 := range alphaChars {
				if keyLen > 3 {

					for _, p4 := range alphaChars {

						key := strings.Builder{}
						key.WriteRune(p1)
						key.WriteRune(p2)
						key.WriteRune(p3)
						key.WriteRune(p4)
						slot := rsb.SlotFromString(key.String())
						slots[int(slot)] = append(slots[int(slot)], key.String())
					}
				} else {
					key := strings.Builder{}
					key.WriteRune(p1)
					key.WriteRune(p2)
					key.WriteRune(p3)
					slot := rsb.SlotFromString(key.String())
					slots[int(slot)] = append(slots[int(slot)], key.String())
				}
			}
		}
	}

	for ps, s := range slots {
		fmt.Printf("Slot: %d\tKeys: %v\n", ps, s)
	}
}
