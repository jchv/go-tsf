package tsf

import "unsafe"

// ITfFnAdviseText COM interface.
type ITfFnAdviseText struct {
	IUnknown
}

// ITfFnAdviseTextVtbl COM interface vtable.
type ITfFnAdviseTextVtbl struct {
	IUnknownVtbl
}

// ITfFnAdviseText returns the ITfFnAdviseText vtable.
func (obj *ITfFnAdviseText) ITfFnAdviseText() *ITfFnAdviseTextVtbl {
	return (*ITfFnAdviseTextVtbl)(unsafe.Pointer(obj.Vtbl))
}
