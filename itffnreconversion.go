package tsf

import "unsafe"

// ITfFnReconversion COM interface.
type ITfFnReconversion struct {
	IUnknown
}

// ITfFnReconversionVtbl COM interface vtable.
type ITfFnReconversionVtbl struct {
	IUnknownVtbl
}

// ITfFnReconversion returns the ITfFnReconversion vtable.
func (obj *ITfFnReconversion) ITfFnReconversion() *ITfFnReconversionVtbl {
	return (*ITfFnReconversionVtbl)(unsafe.Pointer(obj.Vtbl))
}
