package ebpfutil

type BPFProgram struct {
	// ID of the BPF program
	ID uint32
	// File Descriptor of the BPF program
	FD uint32
	// Available information: https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L5840
	Info *BPFProgInfo
}

// GetAllStats returns an array of the attached BPF programs and its avaialble information
func GetAllStats() ([]BPFProgram, error) {
	progs, err := ProgramsID()
	if err != nil {
		return nil, err
	}
	bpfprogs := make([]BPFProgram, len(progs))
	for i, progID := range progs {
		fd, err := GetProgFileDescriptor(progID)
		if err != nil {
			return nil, err
		}
		info, err := GetInfoByFD(fd)
		if err != nil {
			return nil, err
		}
		bpfprogs[i] = BPFProgram{
			ID:   progID,
			FD:   fd,
			Info: info,
		}
	}
	return bpfprogs, nil
}
