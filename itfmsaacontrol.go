package tsf

import "unsafe"

// ITfMSAAControl COM interface.
type ITfMSAAControl struct {
	IUnknown
}

// ITfMSAAControlVtbl COM interface vtable.
type ITfMSAAControlVtbl struct {
	IUnknownVtbl
}

// ITfMSAAControl returns the ITfMSAAControl vtable.
func (obj *ITfMSAAControl) ITfMSAAControl() *ITfMSAAControlVtbl {
	return (*ITfMSAAControlVtbl)(unsafe.Pointer(obj.Vtbl))
}
