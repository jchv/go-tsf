package tsf

import "unsafe"

// ITfReverseConversion COM interface.
type ITfReverseConversion struct {
	IUnknown
}

// ITfReverseConversionVtbl COM interface vtable.
type ITfReverseConversionVtbl struct {
	IUnknownVtbl
}

// ITfReverseConversion returns the ITfReverseConversion vtable.
func (obj *ITfReverseConversion) ITfReverseConversion() *ITfReverseConversionVtbl {
	return (*ITfReverseConversionVtbl)(unsafe.Pointer(obj.Vtbl))
}
