package tsf

import "unsafe"

// ITfFnConfigure COM interface.
type ITfFnConfigure struct {
	IUnknown
}

// ITfFnConfigureVtbl COM interface vtable.
type ITfFnConfigureVtbl struct {
	IUnknownVtbl
}

// ITfFnConfigure returns the ITfFnConfigure vtable.
func (obj *ITfFnConfigure) ITfFnConfigure() *ITfFnConfigureVtbl {
	return (*ITfFnConfigureVtbl)(unsafe.Pointer(obj.Vtbl))
}
