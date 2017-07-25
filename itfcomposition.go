package tsf

import "unsafe"

// ITfComposition COM interface.
type ITfComposition struct {
	IUnknown
}

// ITfCompositionVtbl COM interface vtable.
type ITfCompositionVtbl struct {
	IUnknownVtbl
}

// ITfComposition returns the ITfComposition vtable.
func (obj *ITfComposition) ITfComposition() *ITfCompositionVtbl {
	return (*ITfCompositionVtbl)(unsafe.Pointer(obj.Vtbl))
}
