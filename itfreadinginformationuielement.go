package tsf

import "unsafe"

// ITfReadingInformationUIElement COM interface.
type ITfReadingInformationUIElement struct {
	IUnknown
}

// ITfReadingInformationUIElementVtbl COM interface vtable.
type ITfReadingInformationUIElementVtbl struct {
	IUnknownVtbl
}

// ITfReadingInformationUIElement returns the ITfReadingInformationUIElement vtable.
func (obj *ITfReadingInformationUIElement) ITfReadingInformationUIElement() *ITfReadingInformationUIElementVtbl {
	return (*ITfReadingInformationUIElementVtbl)(unsafe.Pointer(obj.Vtbl))
}
