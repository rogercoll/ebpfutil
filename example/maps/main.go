package main

import (
	"fmt"
	"log"

	"github.com/rogercoll/ebpfutil/stats"
)

func main() {
	maps, err := stats.BPFMaps()
	if err != nil {
		log.Fatal(err)
	}
	for _, bpfMap := range maps {
		fmt.Printf("Map name: %v\n", string(bpfMap.Info.Name[:]))
	}
}
