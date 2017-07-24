package tsf

import (
	"syscall"
	"unsafe"
)

// TFRange represents the ITfRange COM interface.
type TFRange struct {
	Vtbl *ITfRange
}

// ITfRange COM interface vtable.
type ITfRange struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr

	GetText           uintptr
	SetText           uintptr
	GetFormattedText  uintptr
	GetEmbedded       uintptr
	InsertEmbedded    uintptr
	ShiftStart        uintptr
	ShiftEnd          uintptr
	ShiftStartToRange uintptr
	ShiftEndToRange   uintptr
	ShiftStartRegion  uintptr
	ShiftEndRegion    uintptr
	IsEmpty           uintptr
	Collapse          uintptr
	IsEqualStart      uintptr
	IsEqualEnd        uintptr
	CompareStart      uintptr
	CompareEnd        uintptr
	AdjustForInsert   uintptr
	GetGravity        uintptr
	SetGravity        uintptr
	Clone             uintptr
	GetContext        uintptr
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on
// an object.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms691379(v=vs.85).aspx
func (obj *TFRange) AddRef() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.AddRef,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}

// Release decrements the reference count for an interface on an object.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682317(v=vs.85).aspx
func (obj *TFRange) Release() uint32 {
	ret, _, _ := syscall.Syscall(
		obj.Vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	return uint32(ret)
}
