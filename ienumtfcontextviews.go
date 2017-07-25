package tsf

import "unsafe"

// IEnumTfContextViews represents the IEnumTfContextViews COM interface.
type IEnumTfContextViews struct {
	IUnknown
}

// IEnumTfContextViewsVtbl COM interface vtable.
type IEnumTfContextViewsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfContextViews returns the IEnumTfContextViews vtable.
func (obj *IEnumTfContextViews) IEnumTfContextViews() *IEnumTfContextViewsVtbl {
	return (*IEnumTfContextViewsVtbl)(unsafe.Pointer(obj.Vtbl))
}
