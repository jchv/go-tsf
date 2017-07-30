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

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfContexts) Clone(enum **IEnumTfContexts) error {
	return result(call1(
		uintptr(obj.IEnumTfContexts().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfContexts) Next(count uint32, items **ITfContext, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfContexts().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfContexts) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfContexts().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfContexts) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfContexts().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
