package tsf

import "unsafe"

// ITfContext COM interface.
type ITfContext struct {
	IUnknown
}

// ITfContextVtbl COM interface vtable.
type ITfContextVtbl struct {
	IUnknownVtbl
}

// ITfContext returns the ITfContext vtable.
func (obj *ITfContext) ITfContext() *ITfContextVtbl {
	return (*ITfContextVtbl)(unsafe.Pointer(obj.Vtbl))
}
