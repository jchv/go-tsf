package tsf

import "unsafe"

// ITfActiveLanguageProfileNotifySink COM interface.
type ITfActiveLanguageProfileNotifySink struct {
	IUnknown
}

// ITfActiveLanguageProfileNotifySinkVtbl COM interface vtable.
type ITfActiveLanguageProfileNotifySinkVtbl struct {
	IUnknownVtbl
}

// ITfActiveLanguageProfileNotifySink returns the ITfActiveLanguageProfileNotifySink vtable.
func (obj *ITfActiveLanguageProfileNotifySink) ITfActiveLanguageProfileNotifySink() *ITfActiveLanguageProfileNotifySinkVtbl {
	return (*ITfActiveLanguageProfileNotifySinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
