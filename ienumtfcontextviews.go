package tsf

import "unsafe"

// IEnumTfContextViews represents the IEnumTfContextViews COM interface.
type IEnumTfContextViews struct {
	IUnknown
}

// IEnumTfContextViewsVtbl COM interface vtable.
type IEnumTfContextViewsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfContextViews returns the IEnumTfContextViews vtable.
func (obj *IEnumTfContextViews) IEnumTfContextViews() *IEnumTfContextViewsVtbl {
	return (*IEnumTfContextViewsVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfContextViews) Clone(enum **IEnumTfContextViews) error {
	return result(call1(
		uintptr(obj.IEnumTfContextViews().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfContextViews) Next(count uint32, items **ITfContextView, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfContextViews().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfContextViews) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfContextViews().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfContextViews) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfContextViews().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
