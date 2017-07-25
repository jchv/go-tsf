package tsf

import "unsafe"

// ITfTextInputProcessor COM interface.
type ITfTextInputProcessor struct {
	IUnknown
}

// ITfTextInputProcessorVtbl COM interface vtable.
type ITfTextInputProcessorVtbl struct {
	IUnknownVtbl
}

// ITfTextInputProcessor returns the ITfTextInputProcessor vtable.
func (obj *ITfTextInputProcessor) ITfTextInputProcessor() *ITfTextInputProcessorVtbl {
	return (*ITfTextInputProcessorVtbl)(unsafe.Pointer(obj.Vtbl))
}
