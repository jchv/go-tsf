package tsf

import "unsafe"

// ITfMenu COM interface.
type ITfMenu struct {
	IUnknown
}

// ITfMenuVtbl COM interface vtable.
type ITfMenuVtbl struct {
	IUnknownVtbl
}

// ITfMenu returns the ITfMenu vtable.
func (obj *ITfMenu) ITfMenu() *ITfMenuVtbl {
	return (*ITfMenuVtbl)(unsafe.Pointer(obj.Vtbl))
}
