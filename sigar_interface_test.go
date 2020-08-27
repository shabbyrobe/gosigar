package sigar

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

const invalidPid = 666666

func TestCPU(t *testing.T) {
	cpu := Cpu{}
	err := cpu.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadAverage(t *testing.T) {
	avg := LoadAverage{}
	err := avg.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
}

func TestUptime(t *testing.T) {
	uptime := Uptime{}
	err := uptime.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if uptime.Length <= 0 {
		t.Fatal()
	}
}

func TestMem(t *testing.T) {
	mem := Mem{}
	err := mem.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	if mem.Total <= 0 {
		t.Fatal()
	}
	if (mem.Used + mem.Free) > mem.Total {
		t.Fatal()
	}
}

func TestSwap(t *testing.T) {
	swap := Swap{}
	err := swap.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if (swap.Used + swap.Free) > swap.Total {
		t.Fatal()
	}
}

func TestCPUList(t *testing.T) {
	cpulist := CpuList{}
	err := cpulist.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	nsigar := len(cpulist.List)
	numcpu := runtime.NumCPU()
	if nsigar != numcpu {
		t.Fatal()
	}
}

func TestFileSystemList(t *testing.T) {
	fslist := FileSystemList{}
	err := fslist.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(fslist.List) == 0 {
		t.Fatal()
	}
}

func TestFileSystemUsage(t *testing.T) {
	fsusage := FileSystemUsage{}
	err := fsusage.Get("/")
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	err = fsusage.Get("T O T A L L Y B O G U S")
	if err == nil {
		t.Fatal(err)
	}
}

func TestProcList(t *testing.T) {
	pids := ProcList{}
	err := pids.Get()
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(pids.List) <= 2 {
		t.Fatal()
	}

	err = pids.Get()
	if err != nil {
		t.Fatal(err)
	}

}

func TestProcState(t *testing.T) {
	state := ProcState{}
	err := state.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	if state.State != RunStateRun && state.State != RunStateSleep {
		t.Fatal()
	}
	if filepath.Base(state.Name) != "go" {
		t.Fatal()
	}

	err = state.Get(invalidPid)
	if err == nil {
		t.Fatal(err)
	}
}

func TestProcCPU(t *testing.T) {
	pCpu := ProcCpu{}
	err := pCpu.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	err = pCpu.Get(invalidPid)
	if err == nil {
		t.Fatal(err)
	}
}

func TestProcMem(t *testing.T) {
	mem := ProcMem{}
	err := mem.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	err = mem.Get(invalidPid)
	if err == nil {
		t.Fatal(err)
	}
}

func TestProcTime(t *testing.T) {
	time := ProcTime{}
	err := time.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}

	err = time.Get(invalidPid)
	if err == nil {
		t.Fatal(err)
	}
}

func TestProcArgs(t *testing.T) {
	args := ProcArgs{}
	err := args.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if len(args.List) < 1 {
		t.Fatal()
	}
}

func TestProcExe(t *testing.T) {
	exe := ProcExe{}
	err := exe.Get(os.Getppid())
	if err == ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if filepath.Base(exe.Name) != "go" {
		t.Fatal(exe.Name)
	}
}
