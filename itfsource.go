package tsf

import "unsafe"

// ITfSource COM interface.
type ITfSource struct {
	IUnknown
}

// ITfSourceVtbl COM interface vtable.
type ITfSourceVtbl struct {
	IUnknownVtbl
}

// ITfSource returns the ITfSource vtable.
func (obj *ITfSource) ITfSource() *ITfSourceVtbl {
	return (*ITfSourceVtbl)(unsafe.Pointer(obj.Vtbl))
}
