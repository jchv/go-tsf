package tsf

import "unsafe"

// ITfRange COM interface.
type ITfRange struct {
	IUnknown
}

// ITfRangeVtbl COM interface vtable.
type ITfRangeVtbl struct {
	IUnknownVtbl
}

// ITfRange returns the ITfRange vtable.
func (obj *ITfRange) ITfRange() *ITfRangeVtbl {
	return (*ITfRangeVtbl)(unsafe.Pointer(obj.Vtbl))
}
