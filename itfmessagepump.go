package tsf

import "unsafe"

// ITfMessagePump COM interface.
type ITfMessagePump struct {
	IUnknown
}

// ITfMessagePumpVtbl COM interface vtable.
type ITfMessagePumpVtbl struct {
	IUnknownVtbl
}

// ITfMessagePump returns the ITfMessagePump vtable.
func (obj *ITfMessagePump) ITfMessagePump() *ITfMessagePumpVtbl {
	return (*ITfMessagePumpVtbl)(unsafe.Pointer(obj.Vtbl))
}
