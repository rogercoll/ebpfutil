package stats

import (
	"fmt"
	"sync"

	"github.com/rogercoll/ebpfutil"
	"go.uber.org/multierr"
)

type BPFMap struct {
	// ID of the BPF program
	ID uint32
	// File Descriptor of the BPF program
	FD uint32
	// Available information: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L5840
	Info *ebpfutil.BPFMapInfo
}

type mapResult struct {
	bpfMap BPFMap
	err    error
}

// BPFMaps returns an array of the attached BPF maps and its avaialble information
func BPFMaps() ([]BPFMap, error) {
	maps, err := ebpfutil.MapsID()
	if err != nil {
		return nil, err
	}

	results := make(chan mapResult, len(maps))
	wg := &sync.WaitGroup{}
	wg.Add(len(maps))
	for _, mapID := range maps {
		go func(id uint32) {
			defer wg.Done()
			data := BPFMap{ID: id}
			fd, err := ebpfutil.GetMapFileDescriptor(id)
			if err != nil {
				results <- mapResult{data, err}
				return
			}
			data.FD = fd
			info, err := ebpfutil.GetMapInfoByFD(fd)
			if err != nil {
				results <- mapResult{data, err}
				return
			}
			data.Info = info
			results <- mapResult{data, nil}
		}(mapID)
	}
	wg.Wait()
	close(results)

	var errs error
	var bpfprogs []BPFMap
	for res := range results {
		if res.err != nil {
			errs = multierr.Append(errs, fmt.Errorf("BPF map %d, error: %w", res.bpfMap.ID, res.err))
			continue
		}
		bpfprogs = append(bpfprogs, res.bpfMap)
	}

	return bpfprogs, errs
}
