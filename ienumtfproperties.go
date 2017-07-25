package tsf

import "unsafe"

// IEnumTfProperties represents the IEnumTfProperties COM interface.
type IEnumTfProperties struct {
	IUnknown
}

// IEnumTfPropertiesVtbl COM interface vtable.
type IEnumTfPropertiesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfProperties returns the IEnumTfProperties vtable.
func (obj *IEnumTfProperties) IEnumTfProperties() *IEnumTfPropertiesVtbl {
	return (*IEnumTfPropertiesVtbl)(unsafe.Pointer(obj.Vtbl))
}
