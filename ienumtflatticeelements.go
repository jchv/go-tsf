package tsf

import "unsafe"

// IEnumTfLatticeElements represents the IEnumTfLatticeElements COM interface.
type IEnumTfLatticeElements struct {
	IUnknown
}

// IEnumTfLatticeElementsVtbl COM interface vtable.
type IEnumTfLatticeElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfLatticeElements returns the IEnumTfLatticeElements vtable.
func (obj *IEnumTfLatticeElements) IEnumTfLatticeElements() *IEnumTfLatticeElementsVtbl {
	return (*IEnumTfLatticeElementsVtbl)(unsafe.Pointer(obj.Vtbl))
}
