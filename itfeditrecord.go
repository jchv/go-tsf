package tsf

import "unsafe"

// ITfEditRecord COM interface.
type ITfEditRecord struct {
	IUnknown
}

// ITfEditRecordVtbl COM interface vtable.
type ITfEditRecordVtbl struct {
	IUnknownVtbl
}

// ITfEditRecord returns the ITfEditRecord vtable.
func (obj *ITfEditRecord) ITfEditRecord() *ITfEditRecordVtbl {
	return (*ITfEditRecordVtbl)(unsafe.Pointer(obj.Vtbl))
}
