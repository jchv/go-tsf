package tsf

import "unsafe"

// ITfContextComposition COM interface.
type ITfContextComposition struct {
	IUnknown
}

// ITfContextCompositionVtbl COM interface vtable.
type ITfContextCompositionVtbl struct {
	IUnknownVtbl
}

// ITfContextComposition returns the ITfContextComposition vtable.
func (obj *ITfContextComposition) ITfContextComposition() *ITfContextCompositionVtbl {
	return (*ITfContextCompositionVtbl)(unsafe.Pointer(obj.Vtbl))
}
