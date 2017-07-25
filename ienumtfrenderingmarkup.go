package tsf

import "unsafe"

// IEnumTfRenderingMarkup represents the IEnumTfRenderingMarkup COM interface.
type IEnumTfRenderingMarkup struct {
	IUnknown
}

// IEnumTfRenderingMarkupVtbl COM interface vtable.
type IEnumTfRenderingMarkupVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfRenderingMarkup returns the IEnumTfRenderingMarkup vtable.
func (obj *IEnumTfRenderingMarkup) IEnumTfRenderingMarkup() *IEnumTfRenderingMarkupVtbl {
	return (*IEnumTfRenderingMarkupVtbl)(unsafe.Pointer(obj.Vtbl))
}
