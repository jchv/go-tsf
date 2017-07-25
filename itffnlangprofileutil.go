package tsf

import "unsafe"

// ITfFnLangProfileUtil COM interface.
type ITfFnLangProfileUtil struct {
	IUnknown
}

// ITfFnLangProfileUtilVtbl COM interface vtable.
type ITfFnLangProfileUtilVtbl struct {
	IUnknownVtbl
}

// ITfFnLangProfileUtil returns the ITfFnLangProfileUtil vtable.
func (obj *ITfFnLangProfileUtil) ITfFnLangProfileUtil() *ITfFnLangProfileUtilVtbl {
	return (*ITfFnLangProfileUtilVtbl)(unsafe.Pointer(obj.Vtbl))
}
