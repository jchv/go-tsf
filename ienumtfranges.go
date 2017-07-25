package tsf

import "unsafe"

// IEnumTfRanges represents the IEnumTfRanges COM interface.
type IEnumTfRanges struct {
	IUnknown
}

// IEnumTfRangesVtbl COM interface vtable.
type IEnumTfRangesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfRanges returns the IEnumTfRanges vtable.
func (obj *IEnumTfRanges) IEnumTfRanges() *IEnumTfRangesVtbl {
	return (*IEnumTfRangesVtbl)(unsafe.Pointer(obj.Vtbl))
}
