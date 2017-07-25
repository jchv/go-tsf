package tsf

import "unsafe"

// ITfContextView COM interface.
type ITfContextView struct {
	IUnknown
}

// ITfContextViewVtbl COM interface vtable.
type ITfContextViewVtbl struct {
	IUnknownVtbl
}

// ITfContextView returns the ITfContextView vtable.
func (obj *ITfContextView) ITfContextView() *ITfContextViewVtbl {
	return (*ITfContextViewVtbl)(unsafe.Pointer(obj.Vtbl))
}
