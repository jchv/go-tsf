package tsf

import "unsafe"

// ITextStoreACPServices COM interface.
type ITextStoreACPServices struct {
	IUnknown
}

// ITextStoreACPServicesVtbl COM interface vtable.
type ITextStoreACPServicesVtbl struct {
	IUnknownVtbl
}

// ITextStoreACPServices returns the ITextStoreACPServices vtable.
func (obj *ITextStoreACPServices) ITextStoreACPServices() *ITextStoreACPServicesVtbl {
	return (*ITextStoreACPServicesVtbl)(unsafe.Pointer(obj.Vtbl))
}
