package tsf

import (
	"syscall"
)

// GUID structure used by Win32 and COM.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

// defineGUID provides a syntax similar to the C "DEFINEGUID" macro.
func defineGUID(a uint32, b uint16, c uint16, d0 uint8, d1 uint8, d2 uint8, d3 uint8, d4 uint8, d5 uint8, d6 uint8, d7 uint8) GUID {
	return GUID{
		Data1: a,
		Data2: b,
		Data3: c,
		Data4: [8]byte{d0, d1, d2, d3, d4, d5, d6, d7},
	}
}

func call0(fn, this uintptr) uintptr {
	ret, _, _ := syscall.Syscall(fn, 1, this, 0, 0)
	return ret
}

func call1(fn, this, a1 uintptr) uintptr {
	ret, _, _ := syscall.Syscall(fn, 2, this, a1, 0)
	return ret
}

func call2(fn, this, a1, a2 uintptr) uintptr {
	ret, _, _ := syscall.Syscall(fn, 3, this, a1, a2)
	return ret
}

func call3(fn, this, a1, a2, a3 uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(fn, 4, this, a1, a2, a3, 0, 0)
	return ret
}

func call4(fn, this, a1, a2, a3, a4 uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(fn, 5, this, a1, a2, a3, a4, 0)
	return ret
}

func call5(fn, this, a1, a2, a3, a4, a5 uintptr) uintptr {
	ret, _, _ := syscall.Syscall6(fn, 6, this, a1, a2, a3, a4, a5)
	return ret
}

func call6(fn, this, a1, a2, a3, a4, a5, a6 uintptr) uintptr {
	ret, _, _ := syscall.Syscall9(fn, 7, this, a1, a2, a3, a4, a5, a6, 0, 0)
	return ret
}

func call7(fn, this, a1, a2, a3, a4, a5, a6, a7 uintptr) uintptr {
	ret, _, _ := syscall.Syscall9(fn, 8, this, a1, a2, a3, a4, a5, a6, a7, 0)
	return ret
}

func call8(fn, this, a1, a2, a3, a4, a5, a6, a7, a8 uintptr) uintptr {
	ret, _, _ := syscall.Syscall9(fn, 9, this, a1, a2, a3, a4, a5, a6, a7, a8)
	return ret
}

func result(hresult uintptr) error {
	if hresult != 0x00000000 {
		return COMError{uint32(hresult)}
	}

	return nil
}
