package xf

import (
	"syscall"
	"unsafe"
)

const (
	MSP_SUCCESS = 0
)

func unsafeString(p string) unsafe.Pointer {
	return unsafe.Pointer(&[]byte(p + "\x00")[0])
}

func uintptrToString(cstr uintptr) string {
    if cstr != 0 {
        us := make([]byte, 0, 256 * 1024)
        for p := cstr; ; p += 1 {
            u := *(*byte)(unsafe.Pointer(p))
            if u == 0 {
                return string(us)
            }
            us = append(us, u)
        }
    }
    return ""
}

