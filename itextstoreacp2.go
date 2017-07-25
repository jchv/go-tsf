package tsf

import "unsafe"

// ITextStoreACP2 COM interface.
type ITextStoreACP2 struct {
	IUnknown
}

// ITextStoreACP2Vtbl COM interface vtable.
type ITextStoreACP2Vtbl struct {
	IUnknownVtbl
}

// ITextStoreACP2 returns the ITextStoreACP2 vtable.
func (obj *ITextStoreACP2) ITextStoreACP2() *ITextStoreACP2Vtbl {
	return (*ITextStoreACP2Vtbl)(unsafe.Pointer(obj.Vtbl))
}
