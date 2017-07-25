package tsf

import "unsafe"

// ITfRangeACP COM interface.
type ITfRangeACP struct {
	IUnknown
}

// ITfRangeACPVtbl COM interface vtable.
type ITfRangeACPVtbl struct {
	IUnknownVtbl
}

// ITfRangeACP returns the ITfRangeACP vtable.
func (obj *ITfRangeACP) ITfRangeACP() *ITfRangeACPVtbl {
	return (*ITfRangeACPVtbl)(unsafe.Pointer(obj.Vtbl))
}
