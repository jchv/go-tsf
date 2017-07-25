package tsf

import "unsafe"

// ITfCleanupContextDurationSink COM interface.
type ITfCleanupContextDurationSink struct {
	IUnknown
}

// ITfCleanupContextDurationSinkVtbl COM interface vtable.
type ITfCleanupContextDurationSinkVtbl struct {
	IUnknownVtbl
}

// ITfCleanupContextDurationSink returns the ITfCleanupContextDurationSink vtable.
func (obj *ITfCleanupContextDurationSink) ITfCleanupContextDurationSink() *ITfCleanupContextDurationSinkVtbl {
	return (*ITfCleanupContextDurationSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
