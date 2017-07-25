package tsf

import "unsafe"

// IEnumTfContexts represents the IEnumTfContexts COM interface.
type IEnumTfContexts struct {
	IUnknown
}

// IEnumTfContextsVtbl COM interface vtable.
type IEnumTfContextsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfContexts returns the IEnumTfContexts vtable.
func (obj *IEnumTfContexts) IEnumTfContexts() *IEnumTfContextsVtbl {
	return (*IEnumTfContextsVtbl)(unsafe.Pointer(obj.Vtbl))
}
