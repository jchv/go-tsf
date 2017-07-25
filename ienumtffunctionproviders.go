package tsf

import "unsafe"

// IEnumTfFunctionProviders represents the IEnumTfFunctionProviders COM interface.
type IEnumTfFunctionProviders struct {
	IUnknown
}

// IEnumTfFunctionProvidersVtbl COM interface vtable.
type IEnumTfFunctionProvidersVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfFunctionProviders returns the IEnumTfFunctionProviders vtable.
func (obj *IEnumTfFunctionProviders) IEnumTfFunctionProviders() *IEnumTfFunctionProvidersVtbl {
	return (*IEnumTfFunctionProvidersVtbl)(unsafe.Pointer(obj.Vtbl))
}
