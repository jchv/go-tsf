package tsf

import "unsafe"

// ITfCompositionSink COM interface.
type ITfCompositionSink struct {
	IUnknown
}

// ITfCompositionSinkVtbl COM interface vtable.
type ITfCompositionSinkVtbl struct {
	IUnknownVtbl
}

// ITfCompositionSink returns the ITfCompositionSink vtable.
func (obj *ITfCompositionSink) ITfCompositionSink() *ITfCompositionSinkVtbl {
	return (*ITfCompositionSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
