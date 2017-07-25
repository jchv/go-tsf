package tsf

import "unsafe"

// ITfQueryEmbedded COM interface.
type ITfQueryEmbedded struct {
	IUnknown
}

// ITfQueryEmbeddedVtbl COM interface vtable.
type ITfQueryEmbeddedVtbl struct {
	IUnknownVtbl
}

// ITfQueryEmbedded returns the ITfQueryEmbedded vtable.
func (obj *ITfQueryEmbedded) ITfQueryEmbedded() *ITfQueryEmbeddedVtbl {
	return (*ITfQueryEmbeddedVtbl)(unsafe.Pointer(obj.Vtbl))
}
