package tsf

import "unsafe"

// ITfInputProcessorProfileMgr COM interface.
type ITfInputProcessorProfileMgr struct {
	IUnknown
}

// ITfInputProcessorProfileMgrVtbl COM interface vtable.
type ITfInputProcessorProfileMgrVtbl struct {
	IUnknownVtbl
}

// ITfInputProcessorProfileMgr returns the ITfInputProcessorProfileMgr vtable.
func (obj *ITfInputProcessorProfileMgr) ITfInputProcessorProfileMgr() *ITfInputProcessorProfileMgrVtbl {
	return (*ITfInputProcessorProfileMgrVtbl)(unsafe.Pointer(obj.Vtbl))
}
