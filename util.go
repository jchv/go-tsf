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

// BStr represents Windows BSTRs. These must be freed manually.
type BStr *uint16

// VT represents the type of variant we are dealing with.
type VT uint16

// These are possible VT values.
const (
	VTEmpty          VT = 0x0
	VTNull           VT = 0x1
	VTI2             VT = 0x2
	VTI4             VT = 0x3
	VTR4             VT = 0x4
	VTR8             VT = 0x5
	VTCY             VT = 0x6
	VTDate           VT = 0x7
	VTBStr           VT = 0x8
	VTDispatch       VT = 0x9
	VTError          VT = 0xa
	VTBool           VT = 0xb
	VTVariant        VT = 0xc
	VTUnknown        VT = 0xd
	VTDecimal        VT = 0xe
	VTI1             VT = 0x10
	VTUI1            VT = 0x11
	VTUI2            VT = 0x12
	VTUI4            VT = 0x13
	VTI8             VT = 0x14
	VTUI8            VT = 0x15
	VTInt            VT = 0x16
	VTUint           VT = 0x17
	VTVoid           VT = 0x18
	VTHResult        VT = 0x19
	VTPtr            VT = 0x1a
	VTSafeArray      VT = 0x1b
	VTCArray         VT = 0x1c
	VTUserDefined    VT = 0x1d
	VTLPStr          VT = 0x1e
	VTLPWStr         VT = 0x1f
	VTRecord         VT = 0x24
	VTIntPtr         VT = 0x25
	VTUintPtr        VT = 0x26
	VTFileTime       VT = 0x40
	VTBlob           VT = 0x41
	VTStream         VT = 0x42
	VTStorage        VT = 0x43
	VTStreamedObject VT = 0x44
	VTStoredObject   VT = 0x45
	VTBlobObject     VT = 0x46
	VTCF             VT = 0x47
	VTClsID          VT = 0x48
	VTBStrBlob       VT = 0xfff
	VTVector         VT = 0x1000
	VTArray          VT = 0x2000
	VTByRef          VT = 0x4000
	VTReserved       VT = 0x8000
	VTIllegal        VT = 0xffff
	VTIllegalMasked  VT = 0xfff
	VTTypeMask       VT = 0xfff
)

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
