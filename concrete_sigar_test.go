package sigar_test

import (
	"runtime"
	"testing"
	"time"

	sigar "github.com/cloudfoundry/gosigar"
)

// It immediately makes first CPU usage available even though it's not very accurate
func TestFirstCPUUsage(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar

	samplesCh, stop := concreteSigar.CollectCpuStats(500 * time.Millisecond)
	defer func() {
		stop <- struct{}{}
	}()

	firstValue := <-samplesCh
	if firstValue.User <= 0 {
		t.Fatal()
	}
}

func TestCPUUsageDelta(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar

	samplesCh, stop := concreteSigar.CollectCpuStats(500 * time.Millisecond)
	defer func() {
		stop <- struct{}{}
	}()

	firstValue := <-samplesCh
	secondValue := <-samplesCh

	if secondValue.User >= firstValue.User {
		t.Fatal()
	}
}

func TestItDoesNotBlock(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar

	_, stop := concreteSigar.CollectCpuStats(10 * time.Millisecond)

	// Sleep long enough for samplesCh to fill at least 2 values
	time.Sleep(20 * time.Millisecond)

	c := time.After(1 * time.Second)
	select {
	case stop <- struct{}{}:
	case <-c:
		t.Fatal()
	}

	// If CollectCpuStats blocks it will never get here
}

func TestGetLoadAverage(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar

	avg, err := concreteSigar.GetLoadAverage()
	if err == sigar.ErrNotImplemented {
		t.Skip("Not implemented on " + runtime.GOOS)
	}
	if err != nil {
		t.Fatal(err)
	}
	if avg.One == 0 {
		t.Fatal()
	}
	if avg.Five == 0 {
		t.Fatal()
	}
	if avg.Fifteen == 0 {
		t.Fatal()
	}
}

func TestGetMem(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar
	mem, err := concreteSigar.GetMem()
	if err != nil {
		t.Fatal(err)
	}
	if mem.Total <= 0 {
		t.Fatal()
	}
	if mem.Used+mem.Free > mem.Total {
		t.Fatal()
	}
}

func TestGetSwap(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar
	swap, err := concreteSigar.GetSwap()
	if err != nil {
		t.Fatal(err)
	}
	if swap.Used+swap.Free > swap.Total {
		t.Fatal()
	}
}

func TestGetSwap2(t *testing.T) {
	var concreteSigar sigar.ConcreteSigar
	fsusage, err := concreteSigar.GetFileSystemUsage("/")
	if err != nil {
		t.Fatal(err)
	}
	if fsusage.Total == 0 {
		t.Fatal()
	}

	fsusage, err = concreteSigar.GetFileSystemUsage("T O T A L L Y B O G U S")
	if err == nil {
		t.Fatal()
	}
	if fsusage.Total != 0 {
		t.Fatal()
	}
}
