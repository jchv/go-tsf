package tsf

import "unsafe"

// ITfMouseTracker COM interface.
type ITfMouseTracker struct {
	IUnknown
}

// ITfMouseTrackerVtbl COM interface vtable.
type ITfMouseTrackerVtbl struct {
	IUnknownVtbl
}

// ITfMouseTracker returns the ITfMouseTracker vtable.
func (obj *ITfMouseTracker) ITfMouseTracker() *ITfMouseTrackerVtbl {
	return (*ITfMouseTrackerVtbl)(unsafe.Pointer(obj.Vtbl))
}
