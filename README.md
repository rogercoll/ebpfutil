# eBPFutil

Retrieves basic information of the pinned BPF programs running in the host. This pacakge **cannot** be used to load, attach, link or unload BPF programs, it can only be used to gather BPF stats.

## Usage

For example, to monitor the total number of BPF programs attached to the system:

```Go
package main

import (
	"fmt"
	"log"
	"github.com/rogercoll/ebpfutil"
)

func main() {
	progs, err := ebpfutil.ProgramsID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Total number of BPF programs: %d\n", len(progs))
}
```

The whole BPF programs information can be gather with `GetAllStats()`, the returned information contains the same fields as the internal [kernel structure](https://elixir.bootlin.com/linux/latest/source/include/uapi/linux/bpf.h#L5840) and the file descriptor. For example, it contains the program's name, owner's ID, etc:


```Go
package main

import (
	"fmt"
	"log"
	"github.com/rogercoll/ebpfutil"
)

func main() {
	stats, err := ebpfutil.GetAllStats()
	if err != nil {
		log.Fatal(err)
	}
	for _, stat := range stats {
		fmt.Printf("Program name: %s, Owner ID: %d\n", string(stat.Info.Name[:]), stat.Info.CreatedByUid)
	}
}
```
