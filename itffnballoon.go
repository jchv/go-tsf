package tsf

import "unsafe"

// ITfFnBalloon COM interface.
type ITfFnBalloon struct {
	IUnknown
}

// ITfFnBalloonVtbl COM interface vtable.
type ITfFnBalloonVtbl struct {
	IUnknownVtbl
}

// ITfFnBalloon returns the ITfFnBalloon vtable.
func (obj *ITfFnBalloon) ITfFnBalloon() *ITfFnBalloonVtbl {
	return (*ITfFnBalloonVtbl)(unsafe.Pointer(obj.Vtbl))
}
