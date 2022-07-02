package ebpfutil

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

type bpfGetFDByID struct {
	id    uint32
	next  uint32
	flags uint32
}

// GetProgramFileDescriptor return the file descirptor of a given BPF program ID
// https://elixir.bootlin.com/linux/latest/source/tools/lib/bpf/bpf.c#L1090
func GetProgFileDescriptor(id uint32) (uint32, error) {
	attr := bpfGetFDByID{
		id: id,
	}

	ret, _, err := unix.Syscall(
		unix.SYS_BPF,
		unix.BPF_PROG_GET_FD_BY_ID,
		uintptr(unsafe.Pointer(&attr)),
		unsafe.Sizeof(attr),
	)

	if err != 0 {
		return 0, err
	}

	return uint32(ret), nil
}
