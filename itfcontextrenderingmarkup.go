package tsf

import "unsafe"

// ITfContextRenderingMarkup COM interface.
type ITfContextRenderingMarkup struct {
	IUnknown
}

// ITfContextRenderingMarkupVtbl COM interface vtable.
type ITfContextRenderingMarkupVtbl struct {
	IUnknownVtbl
}

// ITfContextRenderingMarkup returns the ITfContextRenderingMarkup vtable.
func (obj *ITfContextRenderingMarkup) ITfContextRenderingMarkup() *ITfContextRenderingMarkupVtbl {
	return (*ITfContextRenderingMarkupVtbl)(unsafe.Pointer(obj.Vtbl))
}
