package tsf

import "unsafe"

// ITfCompartment COM interface.
type ITfCompartment struct {
	IUnknown
}

// ITfCompartmentVtbl COM interface vtable.
type ITfCompartmentVtbl struct {
	IUnknownVtbl
}

// ITfCompartment returns the ITfCompartment vtable.
func (obj *ITfCompartment) ITfCompartment() *ITfCompartmentVtbl {
	return (*ITfCompartmentVtbl)(unsafe.Pointer(obj.Vtbl))
}
