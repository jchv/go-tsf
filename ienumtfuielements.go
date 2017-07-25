package tsf

import "unsafe"

// IEnumTfUIElements represents the IEnumTfUIElements COM interface.
type IEnumTfUIElements struct {
	IUnknown
}

// IEnumTfUIElementsVtbl COM interface vtable.
type IEnumTfUIElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfUIElements returns the IEnumTfUIElements vtable.
func (obj *IEnumTfUIElements) IEnumTfUIElements() *IEnumTfUIElementsVtbl {
	return (*IEnumTfUIElementsVtbl)(unsafe.Pointer(obj.Vtbl))
}
