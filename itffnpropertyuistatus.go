package tsf

import "unsafe"

// ITfFnPropertyUIStatus COM interface.
type ITfFnPropertyUIStatus struct {
	IUnknown
}

// ITfFnPropertyUIStatusVtbl COM interface vtable.
type ITfFnPropertyUIStatusVtbl struct {
	IUnknownVtbl
}

// ITfFnPropertyUIStatus returns the ITfFnPropertyUIStatus vtable.
func (obj *ITfFnPropertyUIStatus) ITfFnPropertyUIStatus() *ITfFnPropertyUIStatusVtbl {
	return (*ITfFnPropertyUIStatusVtbl)(unsafe.Pointer(obj.Vtbl))
}
