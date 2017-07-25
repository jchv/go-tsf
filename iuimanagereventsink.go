package tsf

import "unsafe"

// IUIManagerEventSink COM interface.
type IUIManagerEventSink struct {
	IUnknown
}

// IUIManagerEventSinkVtbl COM interface vtable.
type IUIManagerEventSinkVtbl struct {
	IUnknownVtbl
}

// IUIManagerEventSink returns the IUIManagerEventSink vtable.
func (obj *IUIManagerEventSink) IUIManagerEventSink() *IUIManagerEventSinkVtbl {
	return (*IUIManagerEventSinkVtbl)(unsafe.Pointer(obj.Vtbl))
}
