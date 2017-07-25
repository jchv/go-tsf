package tsf

import "unsafe"

// ITfClientId COM interface.
type ITfClientId struct {
	IUnknown
}

// ITfClientIdVtbl COM interface vtable.
type ITfClientIdVtbl struct {
	IUnknownVtbl
}

// ITfClientId returns the ITfClientId vtable.
func (obj *ITfClientId) ITfClientId() *ITfClientIdVtbl {
	return (*ITfClientIdVtbl)(unsafe.Pointer(obj.Vtbl))
}
