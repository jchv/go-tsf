package tsf

import "unsafe"

// ITfContextOwner COM interface.
type ITfContextOwner struct {
	IUnknown
}

// ITfContextOwnerVtbl COM interface vtable.
type ITfContextOwnerVtbl struct {
	IUnknownVtbl
}

// ITfContextOwner returns the ITfContextOwner vtable.
func (obj *ITfContextOwner) ITfContextOwner() *ITfContextOwnerVtbl {
	return (*ITfContextOwnerVtbl)(unsafe.Pointer(obj.Vtbl))
}
