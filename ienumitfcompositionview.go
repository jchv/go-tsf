package tsf

import "unsafe"

// IEnumITfCompositionView represents the IEnumITfCompositionView COM interface.
type IEnumITfCompositionView struct {
	IUnknown
}

// IEnumITfCompositionViewVtbl COM interface vtable.
type IEnumITfCompositionViewVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumITfCompositionView returns the IEnumITfCompositionView vtable.
func (obj *IEnumITfCompositionView) IEnumITfCompositionView() *IEnumITfCompositionViewVtbl {
	return (*IEnumITfCompositionViewVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumITfCompositionView) Clone(enum **IEnumITfCompositionView) error {
	return result(call1(
		uintptr(obj.IEnumITfCompositionView().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumITfCompositionView) Next(count uint32, views **ITfCompositionView, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumITfCompositionView().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(views)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumITfCompositionView) Reset() error {
	return result(call0(
		uintptr(obj.IEnumITfCompositionView().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumITfCompositionView) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumITfCompositionView().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
