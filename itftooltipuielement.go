package tsf

import "unsafe"

// ITfToolTipUIElement COM interface.
type ITfToolTipUIElement struct {
	IUnknown
}

// ITfToolTipUIElementVtbl COM interface vtable.
type ITfToolTipUIElementVtbl struct {
	IUnknownVtbl
}

// ITfToolTipUIElement returns the ITfToolTipUIElement vtable.
func (obj *ITfToolTipUIElement) ITfToolTipUIElement() *ITfToolTipUIElementVtbl {
	return (*ITfToolTipUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
