package tsf

import "unsafe"

// IEnumITfCompositionView represents the IEnumITfCompositionView COM interface.
type IEnumITfCompositionView struct {
	IUnknown
}

// IEnumITfCompositionViewVtbl COM interface vtable.
type IEnumITfCompositionViewVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumITfCompositionView returns the IEnumITfCompositionView vtable.
func (obj *IEnumITfCompositionView) IEnumITfCompositionView() *IEnumITfCompositionViewVtbl {
	return (*IEnumITfCompositionViewVtbl)(unsafe.Pointer(obj.Vtbl))
}
