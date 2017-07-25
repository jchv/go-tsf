package tsf

import "unsafe"

// ITfTextInputProcessorEx COM interface.
type ITfTextInputProcessorEx struct {
	IUnknown
}

// ITfTextInputProcessorExVtbl COM interface vtable.
type ITfTextInputProcessorExVtbl struct {
	IUnknownVtbl
}

// ITfTextInputProcessorEx returns the ITfTextInputProcessorEx vtable.
func (obj *ITfTextInputProcessorEx) ITfTextInputProcessorEx() *ITfTextInputProcessorExVtbl {
	return (*ITfTextInputProcessorExVtbl)(unsafe.Pointer(obj.Vtbl))
}
