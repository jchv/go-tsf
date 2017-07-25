package tsf

import "unsafe"

// ITfDisplayAttributeInfo COM interface.
type ITfDisplayAttributeInfo struct {
	IUnknown
}

// ITfDisplayAttributeInfoVtbl COM interface vtable.
type ITfDisplayAttributeInfoVtbl struct {
	IUnknownVtbl
}

// ITfDisplayAttributeInfo returns the ITfDisplayAttributeInfo vtable.
func (obj *ITfDisplayAttributeInfo) ITfDisplayAttributeInfo() *ITfDisplayAttributeInfoVtbl {
	return (*ITfDisplayAttributeInfoVtbl)(unsafe.Pointer(obj.Vtbl))
}
