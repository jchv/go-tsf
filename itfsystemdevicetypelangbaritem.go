package tsf

import "unsafe"

// ITfSystemDeviceTypeLangBarItem COM interface.
type ITfSystemDeviceTypeLangBarItem struct {
	IUnknown
}

// ITfSystemDeviceTypeLangBarItemVtbl COM interface vtable.
type ITfSystemDeviceTypeLangBarItemVtbl struct {
	IUnknownVtbl
}

// ITfSystemDeviceTypeLangBarItem returns the ITfSystemDeviceTypeLangBarItem vtable.
func (obj *ITfSystemDeviceTypeLangBarItem) ITfSystemDeviceTypeLangBarItem() *ITfSystemDeviceTypeLangBarItemVtbl {
	return (*ITfSystemDeviceTypeLangBarItemVtbl)(unsafe.Pointer(obj.Vtbl))
}
