package tsf

import "unsafe"

// ITfFnLMInternal COM interface.
type ITfFnLMInternal struct {
	IUnknown
}

// ITfFnLMInternalVtbl COM interface vtable.
type ITfFnLMInternalVtbl struct {
	IUnknownVtbl
}

// ITfFnLMInternal returns the ITfFnLMInternal vtable.
func (obj *ITfFnLMInternal) ITfFnLMInternal() *ITfFnLMInternalVtbl {
	return (*ITfFnLMInternalVtbl)(unsafe.Pointer(obj.Vtbl))
}
