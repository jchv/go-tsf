package tsf

import "unsafe"

// ITfFunctionProvider COM interface.
type ITfFunctionProvider struct {
	IUnknown
}

// ITfFunctionProviderVtbl COM interface vtable.
type ITfFunctionProviderVtbl struct {
	IUnknownVtbl
}

// ITfFunctionProvider returns the ITfFunctionProvider vtable.
func (obj *ITfFunctionProvider) ITfFunctionProvider() *ITfFunctionProviderVtbl {
	return (*ITfFunctionProviderVtbl)(unsafe.Pointer(obj.Vtbl))
}
