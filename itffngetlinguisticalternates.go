package tsf

import "unsafe"

// ITfFnGetLinguisticAlternates COM interface.
type ITfFnGetLinguisticAlternates struct {
	IUnknown
}

// ITfFnGetLinguisticAlternatesVtbl COM interface vtable.
type ITfFnGetLinguisticAlternatesVtbl struct {
	IUnknownVtbl
}

// ITfFnGetLinguisticAlternates returns the ITfFnGetLinguisticAlternates vtable.
func (obj *ITfFnGetLinguisticAlternates) ITfFnGetLinguisticAlternates() *ITfFnGetLinguisticAlternatesVtbl {
	return (*ITfFnGetLinguisticAlternatesVtbl)(unsafe.Pointer(obj.Vtbl))
}
