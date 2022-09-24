package main

import (
	"fmt"
	"log"

	"github.com/rogercoll/ebpfutil/stats"
)

func main() {
	progs, err := stats.BPFPrograms()
	if err != nil {
		log.Fatal(err)
	}
	for _, prog := range progs {
		info := prog.Info
		//	fmt.Printf("Name: %v\n", string(info.Name))
		fmt.Printf("Program ID: %v, FD: %v, CreatedBy: %v, Name: %v\n", prog.ID, prog.FD, info.CreatedByUid, string(info.Name[:]))
	}
}
