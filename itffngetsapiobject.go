package tsf

import "unsafe"

// ITfFnGetSAPIObject COM interface.
type ITfFnGetSAPIObject struct {
	IUnknown
}

// ITfFnGetSAPIObjectVtbl COM interface vtable.
type ITfFnGetSAPIObjectVtbl struct {
	IUnknownVtbl
}

// ITfFnGetSAPIObject returns the ITfFnGetSAPIObject vtable.
func (obj *ITfFnGetSAPIObject) ITfFnGetSAPIObject() *ITfFnGetSAPIObjectVtbl {
	return (*ITfFnGetSAPIObjectVtbl)(unsafe.Pointer(obj.Vtbl))
}
