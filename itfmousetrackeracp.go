package tsf

import "unsafe"

// ITfMouseTrackerACP COM interface.
type ITfMouseTrackerACP struct {
	IUnknown
}

// ITfMouseTrackerACPVtbl COM interface vtable.
type ITfMouseTrackerACPVtbl struct {
	IUnknownVtbl
}

// ITfMouseTrackerACP returns the ITfMouseTrackerACP vtable.
func (obj *ITfMouseTrackerACP) ITfMouseTrackerACP() *ITfMouseTrackerACPVtbl {
	return (*ITfMouseTrackerACPVtbl)(unsafe.Pointer(obj.Vtbl))
}
