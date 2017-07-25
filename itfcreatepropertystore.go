package tsf

import "unsafe"

// ITfCreatePropertyStore COM interface.
type ITfCreatePropertyStore struct {
	IUnknown
}

// ITfCreatePropertyStoreVtbl COM interface vtable.
type ITfCreatePropertyStoreVtbl struct {
	IUnknownVtbl
}

// ITfCreatePropertyStore returns the ITfCreatePropertyStore vtable.
func (obj *ITfCreatePropertyStore) ITfCreatePropertyStore() *ITfCreatePropertyStoreVtbl {
	return (*ITfCreatePropertyStoreVtbl)(unsafe.Pointer(obj.Vtbl))
}
