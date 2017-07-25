package tsf

import "unsafe"

// ITfPropertyStore COM interface.
type ITfPropertyStore struct {
	IUnknown
}

// ITfPropertyStoreVtbl COM interface vtable.
type ITfPropertyStoreVtbl struct {
	IUnknownVtbl
}

// ITfPropertyStore returns the ITfPropertyStore vtable.
func (obj *ITfPropertyStore) ITfPropertyStore() *ITfPropertyStoreVtbl {
	return (*ITfPropertyStoreVtbl)(unsafe.Pointer(obj.Vtbl))
}
