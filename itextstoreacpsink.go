package tsf

import "unsafe"

// ITextStoreACPSink COM interface.
type ITextStoreACPSink struct {
	IUnknown
}

// ITextStoreACPSinkVtbl COM interface vtable.
type ITextStoreACPSinkVtbl struct {
	IUnknownVtbl
}

// ITextStoreACPSink returns the ITextStoreACPSink vtable.
func (obj *ITextStoreACPSink) ITextStoreACPSink() *ITextStoreACPSinkVtbl {
	return (*ITextStoreACPSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
