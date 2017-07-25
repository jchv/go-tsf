package tsf

import "unsafe"

// ITfContextOwnerCompositionSink COM interface.
type ITfContextOwnerCompositionSink struct {
	IUnknown
}

// ITfContextOwnerCompositionSinkVtbl COM interface vtable.
type ITfContextOwnerCompositionSinkVtbl struct {
	IUnknownVtbl
}

// ITfContextOwnerCompositionSink returns the ITfContextOwnerCompositionSink vtable.
func (obj *ITfContextOwnerCompositionSink) ITfContextOwnerCompositionSink() *ITfContextOwnerCompositionSinkVtbl {
	return (*ITfContextOwnerCompositionSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
