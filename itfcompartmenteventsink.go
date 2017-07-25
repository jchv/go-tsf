package tsf

import "unsafe"

// ITfCompartmentEventSink COM interface.
type ITfCompartmentEventSink struct {
	IUnknown
}

// ITfCompartmentEventSinkVtbl COM interface vtable.
type ITfCompartmentEventSinkVtbl struct {
	IUnknownVtbl
}

// ITfCompartmentEventSink returns the ITfCompartmentEventSink vtable.
func (obj *ITfCompartmentEventSink) ITfCompartmentEventSink() *ITfCompartmentEventSinkVtbl {
	return (*ITfCompartmentEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
