package tsf

import "unsafe"

// ITfCompartmentMgr COM interface.
type ITfCompartmentMgr struct {
	IUnknown
}

// ITfCompartmentMgrVtbl COM interface vtable.
type ITfCompartmentMgrVtbl struct {
	IUnknownVtbl
}

// ITfCompartmentMgr returns the ITfCompartmentMgr vtable.
func (obj *ITfCompartmentMgr) ITfCompartmentMgr() *ITfCompartmentMgrVtbl {
	return (*ITfCompartmentMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
