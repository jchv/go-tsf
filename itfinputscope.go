package tsf

import "unsafe"

// ITfInputScope COM interface.
type ITfInputScope struct {
	IUnknown
}

// ITfInputScopeVtbl COM interface vtable.
type ITfInputScopeVtbl struct {
	IUnknownVtbl
}

// ITfInputScope returns the ITfInputScope vtable.
func (obj *ITfInputScope) ITfInputScope() *ITfInputScopeVtbl {
	return (*ITfInputScopeVtbl)(unsafe.Pointer(obj.Vtbl))
}
