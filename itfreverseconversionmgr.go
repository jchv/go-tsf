package tsf

import "unsafe"

// ITfReverseConversionMgr COM interface.
type ITfReverseConversionMgr struct {
	IUnknown
}

// ITfReverseConversionMgrVtbl COM interface vtable.
type ITfReverseConversionMgrVtbl struct {
	IUnknownVtbl
}

// ITfReverseConversionMgr returns the ITfReverseConversionMgr vtable.
func (obj *ITfReverseConversionMgr) ITfReverseConversionMgr() *ITfReverseConversionMgrVtbl {
	return (*ITfReverseConversionMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
