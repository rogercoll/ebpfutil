package ebpfutil

import "golang.org/x/sys/unix"

// https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L6003
type BPFMapInfo struct {
	MapType               uint32
	ID                    uint32
	KeySize               uint32
	ValueSize             uint32
	MaxEntries            uint32
	MapFlags              uint32
	Name                  [unix.BPF_OBJ_NAME_LEN]byte
	Ifindex               uint32
	BpfVmlinuxValueTypeID uint32
	NetnsDev              uint32
	NetnsIno              uint32
	BtfID                 uint32
	BtfKeyTypeID          uint32
	BtfValueTypeID        uint32
	alignment             uint32
	MapExtra              uint64
}
