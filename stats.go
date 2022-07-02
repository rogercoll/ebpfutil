package ebpfutil

import (
	"fmt"
	"sync"

	"go.uber.org/multierr"
)

type BPFProgram struct {
	// ID of the BPF program
	ID uint32
	// File Descriptor of the BPF program
	FD uint32
	// Available information: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L5840
	Info *BPFProgInfo
}

type result struct {
	program BPFProgram
	err     error
}

// GetAllStats returns an array of the attached BPF programs and its avaialble information
func GetAllStats() ([]BPFProgram, error) {
	progs, err := ProgramsID()
	if err != nil {
		return nil, err
	}

	results := make(chan result, len(progs))
	wg := &sync.WaitGroup{}
	wg.Add(len(progs))
	for _, progID := range progs {
		go func(id uint32) {
			defer wg.Done()
			program := BPFProgram{ID: id}
			fd, err := GetProgFileDescriptor(id)
			if err != nil {
				results <- result{program, err}
				return
			}
			program.FD = fd
			info, err := GetInfoByFD(fd)
			if err != nil {
				results <- result{program, err}
				return
			}
			program.Info = info
			results <- result{program, nil}
		}(progID)
	}
	wg.Wait()
	close(results)

	var errs error
	var bpfprogs []BPFProgram
	for res := range results {
		if res.err != nil {
			errs = multierr.Append(errs, fmt.Errorf("BPF program %d, error: %w", res.program.ID, res.err))
			continue
		}
		bpfprogs = append(bpfprogs, res.program)
	}

	return bpfprogs, errs
}
