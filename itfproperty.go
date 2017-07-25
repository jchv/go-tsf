package tsf

import "unsafe"

// ITfProperty COM interface.
type ITfProperty struct {
	IUnknown
}

// ITfPropertyVtbl COM interface vtable.
type ITfPropertyVtbl struct {
	IUnknownVtbl
}

// ITfProperty returns the ITfProperty vtable.
func (obj *ITfProperty) ITfProperty() *ITfPropertyVtbl {
	return (*ITfPropertyVtbl)(unsafe.Pointer(obj.Vtbl))
}
