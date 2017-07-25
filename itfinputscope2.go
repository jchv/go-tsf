package tsf

import "unsafe"

// ITfInputScope2 COM interface.
type ITfInputScope2 struct {
	IUnknown
}

// ITfInputScope2Vtbl COM interface vtable.
type ITfInputScope2Vtbl struct {
	IUnknownVtbl
}

// ITfInputScope2 returns the ITfInputScope2 vtable.
func (obj *ITfInputScope2) ITfInputScope2() *ITfInputScope2Vtbl {
	return (*ITfInputScope2Vtbl)(unsafe.Pointer(obj.Vtbl))
}
