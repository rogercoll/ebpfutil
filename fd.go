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
func getFileDescriptor(structure uintptr, id uint32) (uint32, error) {
	attr := bpfGetFDByID{
		id: id,
	}

	ret, _, err := unix.Syscall(
		unix.SYS_BPF,
		structure,
		uintptr(unsafe.Pointer(&attr)),
		unsafe.Sizeof(attr),
	)

	if err != 0 {
		return 0, err
	}

	return uint32(ret), nil
}

// GetProgramFileDescriptor return the file descirptor of a given BPF program ID
func GetProgFileDescriptor(id uint32) (uint32, error) {
	return getFileDescriptor(unix.BPF_PROG_GET_FD_BY_ID, id)
}

// GetMapFileDescriptor return the file descirptor of a given BPF map ID
func GetMapFileDescriptor(id uint32) (uint32, error) {
	return getFileDescriptor(unix.BPF_MAP_GET_FD_BY_ID, id)
}
