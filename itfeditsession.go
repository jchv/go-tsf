package tsf

import "unsafe"

// ITfEditSession COM interface.
type ITfEditSession struct {
	IUnknown
}

// ITfEditSessionVtbl COM interface vtable.
type ITfEditSessionVtbl struct {
	IUnknownVtbl
}

// ITfEditSession returns the ITfEditSession vtable.
func (obj *ITfEditSession) ITfEditSession() *ITfEditSessionVtbl {
	return (*ITfEditSessionVtbl)(unsafe.Pointer(obj.Vtbl))
}
