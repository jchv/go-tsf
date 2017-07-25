package tsf

import "unsafe"

// ITfCompositionView COM interface.
type ITfCompositionView struct {
	IUnknown
}

// ITfCompositionViewVtbl COM interface vtable.
type ITfCompositionViewVtbl struct {
	IUnknownVtbl
}

// ITfCompositionView returns the ITfCompositionView vtable.
func (obj *ITfCompositionView) ITfCompositionView() *ITfCompositionViewVtbl {
	return (*ITfCompositionViewVtbl)(unsafe.Pointer(obj.Vtbl))
}
