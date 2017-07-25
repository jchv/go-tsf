package tsf

import "unsafe"

// ITfFnShowHelp COM interface.
type ITfFnShowHelp struct {
	IUnknown
}

// ITfFnShowHelpVtbl COM interface vtable.
type ITfFnShowHelpVtbl struct {
	IUnknownVtbl
}

// ITfFnShowHelp returns the ITfFnShowHelp vtable.
func (obj *ITfFnShowHelp) ITfFnShowHelp() *ITfFnShowHelpVtbl {
	return (*ITfFnShowHelpVtbl)(unsafe.Pointer(obj.Vtbl))
}
