package tsf

import "unsafe"

// ITfFnConfigureRegisterWord COM interface.
type ITfFnConfigureRegisterWord struct {
	IUnknown
}

// ITfFnConfigureRegisterWordVtbl COM interface vtable.
type ITfFnConfigureRegisterWordVtbl struct {
	IUnknownVtbl
}

// ITfFnConfigureRegisterWord returns the ITfFnConfigureRegisterWord vtable.
func (obj *ITfFnConfigureRegisterWord) ITfFnConfigureRegisterWord() *ITfFnConfigureRegisterWordVtbl {
	return (*ITfFnConfigureRegisterWordVtbl)(unsafe.Pointer(obj.Vtbl))
}
