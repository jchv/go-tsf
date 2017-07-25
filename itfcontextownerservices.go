package tsf

import "unsafe"

// ITfContextOwnerServices COM interface.
type ITfContextOwnerServices struct {
	IUnknown
}

// ITfContextOwnerServicesVtbl COM interface vtable.
type ITfContextOwnerServicesVtbl struct {
	IUnknownVtbl
}

// ITfContextOwnerServices returns the ITfContextOwnerServices vtable.
func (obj *ITfContextOwnerServices) ITfContextOwnerServices() *ITfContextOwnerServicesVtbl {
	return (*ITfContextOwnerServicesVtbl)(unsafe.Pointer(obj.Vtbl))
}
