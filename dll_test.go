// +build windows

package xf

import (
	"syscall"
	"testing"
)

func TestDll(t *testing.T) {
	dlldemo := syscall.NewLazyDLL("dlldemo.dll")
	procfn := dlldemo.NewProc("fndlldemo")

	r, _, _ := procfn.Call()
	t.Logf("ret: %d", r)
}

func TestXunfeiDll(t *testing.T) {
	dlldemo := syscall.NewLazyDLL("xunfei.dll")
	procfn := dlldemo.NewProc("fndemo")

	r, _, _ := procfn.Call()
	t.Logf("ret: %d", r)
}
