package tsf

import "unsafe"

// ITfFnLMProcessor COM interface.
type ITfFnLMProcessor struct {
	IUnknown
}

// ITfFnLMProcessorVtbl COM interface vtable.
type ITfFnLMProcessorVtbl struct {
	IUnknownVtbl
}

// ITfFnLMProcessor returns the ITfFnLMProcessor vtable.
func (obj *ITfFnLMProcessor) ITfFnLMProcessor() *ITfFnLMProcessorVtbl {
	return (*ITfFnLMProcessorVtbl)(unsafe.Pointer(obj.Vtbl))
}
