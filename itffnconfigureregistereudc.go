package tsf

import "unsafe"

// ITfFnConfigureRegisterEudc COM interface.
type ITfFnConfigureRegisterEudc struct {
	IUnknown
}

// ITfFnConfigureRegisterEudcVtbl COM interface vtable.
type ITfFnConfigureRegisterEudcVtbl struct {
	IUnknownVtbl
}

// ITfFnConfigureRegisterEudc returns the ITfFnConfigureRegisterEudc vtable.
func (obj *ITfFnConfigureRegisterEudc) ITfFnConfigureRegisterEudc() *ITfFnConfigureRegisterEudcVtbl {
	return (*ITfFnConfigureRegisterEudcVtbl)(unsafe.Pointer(obj.Vtbl))
}
