package tsf

import "unsafe"

// ITfLanguageProfileNotifySink COM interface.
type ITfLanguageProfileNotifySink struct {
	IUnknown
}

// ITfLanguageProfileNotifySinkVtbl COM interface vtable.
type ITfLanguageProfileNotifySinkVtbl struct {
	IUnknownVtbl
}

// ITfLanguageProfileNotifySink returns the ITfLanguageProfileNotifySink vtable.
func (obj *ITfLanguageProfileNotifySink) ITfLanguageProfileNotifySink() *ITfLanguageProfileNotifySinkVtbl {
	return (*ITfLanguageProfileNotifySinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
