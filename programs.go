package ebpfutil

import (
	"errors"
	"unsafe"

	"golang.org/x/sys/unix"
)

/*
struct { anonymous struct used by BPF_*_GET_*_ID
	union {
		__u32           start_id;
		__u32           prog_id;
		__u32           map_id;
		__u32           btf_id;
		__u32           link_id;
																				};
     __u32           next_id;
     __u32           open_flags;
};
*/

type bpfGetId struct {
	ID        uint32
	NextID    uint32
	OpenFlags uint32
}

func progGetNextId(prev uint32) (uint32, error) {
	bgi := bpfGetId{ID: prev}

	_, _, err := unix.Syscall(
		unix.SYS_BPF,
		unix.BPF_PROG_GET_NEXT_ID,
		uintptr(unsafe.Pointer(&bgi)),
		unsafe.Sizeof(bgi),
	)
	if err != 0 {
		return 0, err
	}

	return bgi.NextID, nil
}

func ProgramsID() ([]uint32, error) {
	//TODO: check if kernel > 3.14
	var progs []uint32
	var prev uint32 // inits to 0
	for {
		id, err := progGetNextId(prev)
		if err != nil {
			if errors.Is(err, unix.ENOENT) {
				// all BPF programms scanned
				break
			}
			return nil, err
		}

		progs = append(progs, id)
		prev = id
	}

	return progs, nil
}
