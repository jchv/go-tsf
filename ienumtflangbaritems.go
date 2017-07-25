package tsf

import "unsafe"

// IEnumTfLangBarItems represents the IEnumTfLangBarItems COM interface.
type IEnumTfLangBarItems struct {
	IUnknown
}

// IEnumTfLangBarItemsVtbl COM interface vtable.
type IEnumTfLangBarItemsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfLangBarItems returns the IEnumTfLangBarItems vtable.
func (obj *IEnumTfLangBarItems) IEnumTfLangBarItems() *IEnumTfLangBarItemsVtbl {
	return (*IEnumTfLangBarItemsVtbl)(unsafe.Pointer(obj.Vtbl))
}
