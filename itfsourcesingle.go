package tsf

import "unsafe"

// ITfSourceSingle COM interface.
type ITfSourceSingle struct {
	IUnknown
}

// ITfSourceSingleVtbl COM interface vtable.
type ITfSourceSingleVtbl struct {
	IUnknownVtbl
}

// ITfSourceSingle returns the ITfSourceSingle vtable.
func (obj *ITfSourceSingle) ITfSourceSingle() *ITfSourceSingleVtbl {
	return (*ITfSourceSingleVtbl)(unsafe.Pointer(obj.Vtbl))
}
