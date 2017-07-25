package tsf

import "unsafe"

// ITfReverseConversionList COM interface.
type ITfReverseConversionList struct {
	IUnknown
}

// ITfReverseConversionListVtbl COM interface vtable.
type ITfReverseConversionListVtbl struct {
	IUnknownVtbl
}

// ITfReverseConversionList returns the ITfReverseConversionList vtable.
func (obj *ITfReverseConversionList) ITfReverseConversionList() *ITfReverseConversionListVtbl {
	return (*ITfReverseConversionListVtbl)(unsafe.Pointer(obj.Vtbl))
}
