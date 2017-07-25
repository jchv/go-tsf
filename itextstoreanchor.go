package tsf

import "unsafe"

// ITextStoreAnchor COM interface.
type ITextStoreAnchor struct {
	IUnknown
}

// ITextStoreAnchorVtbl COM interface vtable.
type ITextStoreAnchorVtbl struct {
	IUnknownVtbl
}

// ITextStoreAnchor returns the ITextStoreAnchor vtable.
func (obj *ITextStoreAnchor) ITextStoreAnchor() *ITextStoreAnchorVtbl {
	return (*ITextStoreAnchorVtbl)(unsafe.Pointer(obj.Vtbl))
}
