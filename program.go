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
