package tsf

import "unsafe"

// ITfTransitoryExtensionUIElement COM interface.
type ITfTransitoryExtensionUIElement struct {
	IUnknown
}

// ITfTransitoryExtensionUIElementVtbl COM interface vtable.
type ITfTransitoryExtensionUIElementVtbl struct {
	IUnknownVtbl
}

// ITfTransitoryExtensionUIElement returns the ITfTransitoryExtensionUIElement vtable.
func (obj *ITfTransitoryExtensionUIElement) ITfTransitoryExtensionUIElement() *ITfTransitoryExtensionUIElementVtbl {
	return (*ITfTransitoryExtensionUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
