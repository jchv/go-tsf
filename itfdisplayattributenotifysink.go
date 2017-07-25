package tsf

import "unsafe"

// ITfDisplayAttributeNotifySink COM interface.
type ITfDisplayAttributeNotifySink struct {
	IUnknown
}

// ITfDisplayAttributeNotifySinkVtbl COM interface vtable.
type ITfDisplayAttributeNotifySinkVtbl struct {
	IUnknownVtbl
}

// ITfDisplayAttributeNotifySink returns the ITfDisplayAttributeNotifySink vtable.
func (obj *ITfDisplayAttributeNotifySink) ITfDisplayAttributeNotifySink() *ITfDisplayAttributeNotifySinkVtbl {
	return (*ITfDisplayAttributeNotifySinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
