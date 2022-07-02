package main

import (
	"fmt"
	"log"

	"github.com/rogercoll/ebpfutil"
)

func main() {
	progs, err := ebpfutil.ProgramsID()
	if err != nil {
		log.Fatal(err)
	}
	for _, progID := range progs {
		fd, err := ebpfutil.GetProgFileDescriptor(progID)
		if err != nil {
			log.Fatal(err)
		}
		info, err := ebpfutil.GetInfoByFD(fd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Program ID: %v, FD: %v, CreatedBy: %v, Name: %v\n", progID, fd, info.CreatedByUid, string(info.Name[:]))
	}
}
