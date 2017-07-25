package tsf

import "unsafe"

// ITfFnGetPreferredTouchKeyboardLayout COM interface.
type ITfFnGetPreferredTouchKeyboardLayout struct {
	IUnknown
}

// ITfFnGetPreferredTouchKeyboardLayoutVtbl COM interface vtable.
type ITfFnGetPreferredTouchKeyboardLayoutVtbl struct {
	IUnknownVtbl
}

// ITfFnGetPreferredTouchKeyboardLayout returns the ITfFnGetPreferredTouchKeyboardLayout vtable.
func (obj *ITfFnGetPreferredTouchKeyboardLayout) ITfFnGetPreferredTouchKeyboardLayout() *ITfFnGetPreferredTouchKeyboardLayoutVtbl {
	return (*ITfFnGetPreferredTouchKeyboardLayoutVtbl)(unsafe.Pointer(obj.Vtbl))
}
