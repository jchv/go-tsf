package tsf

import "unsafe"

// ITfTransitoryExtensionSink COM interface.
type ITfTransitoryExtensionSink struct {
	IUnknown
}

// ITfTransitoryExtensionSinkVtbl COM interface vtable.
type ITfTransitoryExtensionSinkVtbl struct {
	IUnknownVtbl
}

// ITfTransitoryExtensionSink returns the ITfTransitoryExtensionSink vtable.
func (obj *ITfTransitoryExtensionSink) ITfTransitoryExtensionSink() *ITfTransitoryExtensionSinkVtbl {
	return (*ITfTransitoryExtensionSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
