package tsf

import "unsafe"

// ITfPersistentPropertyLoaderACP COM interface.
type ITfPersistentPropertyLoaderACP struct {
	IUnknown
}

// ITfPersistentPropertyLoaderACPVtbl COM interface vtable.
type ITfPersistentPropertyLoaderACPVtbl struct {
	IUnknownVtbl
}

// ITfPersistentPropertyLoaderACP returns the ITfPersistentPropertyLoaderACP vtable.
func (obj *ITfPersistentPropertyLoaderACP) ITfPersistentPropertyLoaderACP() *ITfPersistentPropertyLoaderACPVtbl {
	return (*ITfPersistentPropertyLoaderACPVtbl)(unsafe.Pointer(obj.Vtbl))
}
