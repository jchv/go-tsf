package tsf

import "unsafe"

// IEnumTfPropertyValue represents the IEnumTfPropertyValue COM interface.
type IEnumTfPropertyValue struct {
	IUnknown
}

// IEnumTfPropertyValueVtbl COM interface vtable.
type IEnumTfPropertyValueVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfPropertyValue returns the IEnumTfPropertyValue vtable.
func (obj *IEnumTfPropertyValue) IEnumTfPropertyValue() *IEnumTfPropertyValueVtbl {
	return (*IEnumTfPropertyValueVtbl)(unsafe.Pointer(obj.Vtbl))
}
