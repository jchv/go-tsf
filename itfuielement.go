package tsf

import "unsafe"

// ITfUIElement COM interface.
type ITfUIElement struct {
	IUnknown
}

// ITfUIElementVtbl COM interface vtable.
type ITfUIElementVtbl struct {
	IUnknownVtbl
}

// ITfUIElement returns the ITfUIElement vtable.
func (obj *ITfUIElement) ITfUIElement() *ITfUIElementVtbl {
	return (*ITfUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
