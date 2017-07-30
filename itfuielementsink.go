package tsf

import "unsafe"

// ITfUIElementSink COM interface.
type ITfUIElementSink struct {
	IUnknown
}

// ITfUIElementSinkVtbl COM interface vtable.
type ITfUIElementSinkVtbl struct {
	IUnknownVtbl
	BeginUIElement  uintptr
	UpdateUIElement uintptr
	EndUIElement    uintptr
}

// ITfUIElementSink returns the ITfUIElementSink vtable.
func (obj *ITfUIElementSink) ITfUIElementSink() *ITfUIElementSinkVtbl {
	return (*ITfUIElementSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}

// BeginUIElement is called when the UIElement started. This sink can let the
// textservice to draw or not to draw the UI element.
func (obj *ITfUIElementSink) BeginUIElement(elementID uint32, show *uint32) error {
	return result(call2(
		uintptr(obj.ITfUIElementSink().BeginUIElement),
		uintptr(unsafe.Pointer(obj)),
		uintptr(elementID),
		uintptr(unsafe.Pointer(show)),
	))
}

// UpdateUIElement is called when the contents of the UIElement is updated.
func (obj *ITfUIElementSink) UpdateUIElement(elementID uint32) error {
	return result(call1(
		uintptr(obj.ITfUIElementSink().UpdateUIElement),
		uintptr(unsafe.Pointer(obj)),
		uintptr(elementID),
	))
}

// EndUIElement is called when the UIElement is finished.
func (obj *ITfUIElementSink) EndUIElement(elementID uint32) error {
	return result(call1(
		uintptr(obj.ITfUIElementSink().EndUIElement),
		uintptr(unsafe.Pointer(obj)),
		uintptr(elementID),
	))
}
