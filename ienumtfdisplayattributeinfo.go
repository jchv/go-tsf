package tsf

import "unsafe"

// IEnumTfDisplayAttributeInfo represents the IEnumTfDisplayAttributeInfo COM interface.
type IEnumTfDisplayAttributeInfo struct {
	IUnknown
}

// IEnumTfDisplayAttributeInfoVtbl COM interface vtable.
type IEnumTfDisplayAttributeInfoVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfDisplayAttributeInfo returns the IEnumTfDisplayAttributeInfo vtable.
func (obj *IEnumTfDisplayAttributeInfo) IEnumTfDisplayAttributeInfo() *IEnumTfDisplayAttributeInfoVtbl {
	return (*IEnumTfDisplayAttributeInfoVtbl)(unsafe.Pointer(obj.Vtbl))
}
