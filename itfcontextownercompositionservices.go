package tsf

import "unsafe"

// ITfContextOwnerCompositionServices COM interface.
type ITfContextOwnerCompositionServices struct {
	IUnknown
}

// ITfContextOwnerCompositionServicesVtbl COM interface vtable.
type ITfContextOwnerCompositionServicesVtbl struct {
	IUnknownVtbl
}

// ITfContextOwnerCompositionServices returns the ITfContextOwnerCompositionServices vtable.
func (obj *ITfContextOwnerCompositionServices) ITfContextOwnerCompositionServices() *ITfContextOwnerCompositionServicesVtbl {
	return (*ITfContextOwnerCompositionServicesVtbl)(unsafe.Pointer(obj.Vtbl))
}
