package tsf

import "unsafe"

// ITfReadOnlyProperty COM interface.
type ITfReadOnlyProperty struct {
	IUnknown
}

// ITfReadOnlyPropertyVtbl COM interface vtable.
type ITfReadOnlyPropertyVtbl struct {
	IUnknownVtbl
}

// ITfReadOnlyProperty returns the ITfReadOnlyProperty vtable.
func (obj *ITfReadOnlyProperty) ITfReadOnlyProperty() *ITfReadOnlyPropertyVtbl {
	return (*ITfReadOnlyPropertyVtbl)(unsafe.Pointer(obj.Vtbl))
}
