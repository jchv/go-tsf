package tsf

import "unsafe"

// ITfDisplayAttributeProvider COM interface.
type ITfDisplayAttributeProvider struct {
	IUnknown
}

// ITfDisplayAttributeProviderVtbl COM interface vtable.
type ITfDisplayAttributeProviderVtbl struct {
	IUnknownVtbl
}

// ITfDisplayAttributeProvider returns the ITfDisplayAttributeProvider vtable.
func (obj *ITfDisplayAttributeProvider) ITfDisplayAttributeProvider() *ITfDisplayAttributeProviderVtbl {
	return (*ITfDisplayAttributeProviderVtbl)(unsafe.Pointer(obj.Vtbl))
}
