package tsf

import "unsafe"

// ITfFnPlayBack COM interface.
type ITfFnPlayBack struct {
	IUnknown
}

// ITfFnPlayBackVtbl COM interface vtable.
type ITfFnPlayBackVtbl struct {
	IUnknownVtbl
}

// ITfFnPlayBack returns the ITfFnPlayBack vtable.
func (obj *ITfFnPlayBack) ITfFnPlayBack() *ITfFnPlayBackVtbl {
	return (*ITfFnPlayBackVtbl)(unsafe.Pointer(obj.Vtbl))
}
