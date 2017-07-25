package tsf

import "unsafe"

// ITfCleanupContextSink COM interface.
type ITfCleanupContextSink struct {
	IUnknown
}

// ITfCleanupContextSinkVtbl COM interface vtable.
type ITfCleanupContextSinkVtbl struct {
	IUnknownVtbl
}

// ITfCleanupContextSink returns the ITfCleanupContextSink vtable.
func (obj *ITfCleanupContextSink) ITfCleanupContextSink() *ITfCleanupContextSinkVtbl {
	return (*ITfCleanupContextSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
