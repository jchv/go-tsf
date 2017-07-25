package tsf

import "unsafe"

// ITfFunction COM interface.
type ITfFunction struct {
	IUnknown
}

// ITfFunctionVtbl COM interface vtable.
type ITfFunctionVtbl struct {
	IUnknownVtbl
}

// ITfFunction returns the ITfFunction vtable.
func (obj *ITfFunction) ITfFunction() *ITfFunctionVtbl {
	return (*ITfFunctionVtbl)(unsafe.Pointer(obj.Vtbl))
}
