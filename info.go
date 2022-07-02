package ebpfutil

import (
	"unsafe"

	"golang.org/x/sys/unix"
)

// https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L5840
type BPFProgInfo struct {
	ProgType             uint32
	ID                   uint32
	Tag                  [unix.BPF_TAG_SIZE]byte
	JitedProgLen         uint32
	XlatedProgLen        uint32
	JitedProgInsns       unsafe.Pointer
	XlatedProgInsns      unsafe.Pointer
	LoadTime             uint64 /* ns since boottime */
	CreatedByUid         uint32
	NrMapIds             uint32
	MapIds               unsafe.Pointer
	Name                 [unix.BPF_OBJ_NAME_LEN]byte // since 4.15 067cae47771c
	Ifindex              uint32
	GplCompatible        uint32
	NetnsDev             uint64
	NetnsIno             uint64
	NrJitedKsyms         uint32
	NrJitedFuncLens      uint32
	JitedKsmyms          unsafe.Pointer
	JitedFuncLens        unsafe.Pointer
	BtfID                uint32
	FuncInfoRecSize      uint32
	FuncInfo             unsafe.Pointer
	NrFuncInfo           uint32
	NrLineInfo           uint32
	LineInfo             unsafe.Pointer
	JitedLineInfo        unsafe.Pointer
	NrJitedLineInfo      uint32
	LineInfoRecSize      uint32
	JitedLineInfoRecSize uint32
	NrProgTags           uint32
	ProgTags             unsafe.Pointer
	RunTimeNs            uint64
	RunCnt               uint64
}

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

// GetInfoByFD returns the BPF program information given its file descriptor
func GetInfoByFD(fd uint32) (*BPFProgInfo, error) {
	var info BPFProgInfo
	attr := bpfGetInfoByFD{
		FD:      fd,
		InfoLen: uint32(unsafe.Sizeof(info)),
		Info:    unsafe.Pointer(&info),
	}

	_, _, err := unix.Syscall(
		unix.SYS_BPF,
		unix.BPF_OBJ_GET_INFO_BY_FD,
		uintptr(unsafe.Pointer(&attr)),
		unsafe.Sizeof(attr),
	)
	if err != 0 {
		return nil, err
	}
	return &info, nil
}
