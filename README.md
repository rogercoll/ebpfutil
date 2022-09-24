# eBPFutil

[![Go Reference](https://pkg.go.dev/badge/github.com/rogercoll/ebpfutil.svg)](https://pkg.go.dev/github.com/rogercoll/ebpfutil)

Retrieves basic information of the pinned BPF programs and maps running in the host. This package **cannot** be used to load, attach, link or unload BPF programs, it can only be used to gather BPF stats.

## Package Features

### BPF Programs

`stats.BPFPrograms()` -> Returns an array of all the available BPF programs loaded in
the system as a [BPFProgram structure](https://github.com/rogercoll/ebpfutil/blob/main/stats/programs.go#L11) which contains the ID, FD and
[BPFProgInfo](https://github.com/rogercoll/ebpfutil/blob/8b5366a7bf3d0c9b142849a4b6e2e62d23d243b1/program.go#L10) of the corresponding program. See [examples/programs/main.go](./examples/programs/main.go)

### BPF Maps

`stats.BPFMaps()` -> Returns an array of all the available BPF maps in
the system as a [BPFMap structure](https://github.com/rogercoll/ebpfutil/blob/main/stats/maps.go#L11) which contains the ID, FD and
[BPFMapInfo](https://github.com/rogercoll/ebpfutil/blob/8b5366a7bf3d0c9b142849a4b6e2e62d23d243b1/map.go#L6) of the corresponding map. See [examples/maps/main.go](./examples/maps/main.go)


For additional exported functionalities you check the [public documentation](https://pkg.go.dev/github.com/rogercoll/ebpfutil).

## Usage

Import it into your Go program and use any of the exported functions:

```bash
go get github.com/rogercoll/ebpfutil
```
