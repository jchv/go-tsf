package tsf

import (
	"unsafe"
)

// IUnknown represents the IUnknown COM interface.
type IUnknown struct {
	Vtbl *IUnknownVtbl
}

// IUnknownVtbl COM interface vtable.
type IUnknownVtbl struct {
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

// IUnknown returns the IUnknown vtable.
func (obj *IUnknown) IUnknown() *IUnknownVtbl {
	return (*IUnknownVtbl)(unsafe.Pointer(obj.Vtbl))
}

// QueryInterface retrieves pointers to the supported interfaces on an object.
func (obj *IUnknown) QueryInterface(riid *GUID, object uintptr) error {
	return result(call2(
		obj.Vtbl.QueryInterface,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(object),
	))
}

// AddRef increments the reference count for an interface on an object. This
// method should be called for every new copy of a pointer to an interface on
// an object.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms691379(v=vs.85).aspx
func (obj *IUnknown) AddRef() uint32 {
	return uint32(call0(obj.Vtbl.AddRef, uintptr(unsafe.Pointer(obj))))
}

// Release decrements the reference count for an interface on an object.
//
// https://msdn.microsoft.com/en-us/library/windows/desktop/ms682317(v=vs.85).aspx
func (obj *IUnknown) Release() uint32 {
	return uint32(call0(obj.Vtbl.Release, uintptr(unsafe.Pointer(obj))))
}
