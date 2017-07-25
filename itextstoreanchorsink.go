package tsf

import "unsafe"

// ITextStoreAnchorSink COM interface.
type ITextStoreAnchorSink struct {
	IUnknown
}

// ITextStoreAnchorSinkVtbl COM interface vtable.
type ITextStoreAnchorSinkVtbl struct {
	IUnknownVtbl
}

// ITextStoreAnchorSink returns the ITextStoreAnchorSink vtable.
func (obj *ITextStoreAnchorSink) ITextStoreAnchorSink() *ITextStoreAnchorSinkVtbl {
	return (*ITextStoreAnchorSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
