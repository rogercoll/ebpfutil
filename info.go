package ebpfutil

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

/*
struct {  anonymous struct used by BPF_OBJ_GET_INFO_BY_FD
	__u32		bpf_fd;
	__u32		info_len;
	__aligned_u64	info;
} info;
*/

type bpfGetInfoByFD struct {
	FD      uint32
	InfoLen uint32
	Info    unsafe.Pointer
}

// getInfoByFD returns the BPF information given its file descriptor
func getInfoByFD(result uintptr, resultSize, fd uint32) error {
	attr := bpfGetInfoByFD{
		FD:      fd,
		InfoLen: resultSize,
		Info:    unsafe.Pointer(result),
	}

	_, _, err := unix.Syscall(
		unix.SYS_BPF,
		unix.BPF_OBJ_GET_INFO_BY_FD,
		uintptr(unsafe.Pointer(&attr)),
		unsafe.Sizeof(attr),
	)
	if err != 0 {
		return err
	}
	return nil
}

// GetProgInfoByFD returns the BPF program information given its file descriptor
func GetProgInfoByFD(fd uint32) (*BPFProgInfo, error) {
	info := BPFProgInfo{}
	err := getInfoByFD(uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info)), fd)
	if err != nil {
		return nil, err
	}

	return &info, nil
}

// GetMapInfoByFD returns the BPF map information given its file descriptor
func GetMapInfoByFD(fd uint32) (*BPFMapInfo, error) {
	info := BPFMapInfo{}
	err := getInfoByFD(uintptr(unsafe.Pointer(&info)), uint32(unsafe.Sizeof(info)), fd)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
