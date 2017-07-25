package tsf

import "unsafe"

// ITfLMLattice COM interface.
type ITfLMLattice struct {
	IUnknown
}

// ITfLMLatticeVtbl COM interface vtable.
type ITfLMLatticeVtbl struct {
	IUnknownVtbl
}

// ITfLMLattice returns the ITfLMLattice vtable.
func (obj *ITfLMLattice) ITfLMLattice() *ITfLMLatticeVtbl {
	return (*ITfLMLatticeVtbl)(unsafe.Pointer(obj.Vtbl))
}
