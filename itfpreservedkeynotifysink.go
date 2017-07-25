package tsf

import "unsafe"

// ITfPreservedKeyNotifySink COM interface.
type ITfPreservedKeyNotifySink struct {
	IUnknown
}

// ITfPreservedKeyNotifySinkVtbl COM interface vtable.
type ITfPreservedKeyNotifySinkVtbl struct {
	IUnknownVtbl
}

// ITfPreservedKeyNotifySink returns the ITfPreservedKeyNotifySink vtable.
func (obj *ITfPreservedKeyNotifySink) ITfPreservedKeyNotifySink() *ITfPreservedKeyNotifySinkVtbl {
	return (*ITfPreservedKeyNotifySinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
