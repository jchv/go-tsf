package tsf

import "unsafe"

// ITfLangBarItemBalloon COM interface.
type ITfLangBarItemBalloon struct {
	IUnknown
}

// ITfLangBarItemBalloonVtbl COM interface vtable.
type ITfLangBarItemBalloonVtbl struct {
	IUnknownVtbl
}

// ITfLangBarItemBalloon returns the ITfLangBarItemBalloon vtable.
func (obj *ITfLangBarItemBalloon) ITfLangBarItemBalloon() *ITfLangBarItemBalloonVtbl {
	return (*ITfLangBarItemBalloonVtbl)(unsafe.Pointer(obj.Vtbl))
}
