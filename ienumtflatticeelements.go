package tsf

import "unsafe"

// IEnumTfLatticeElements represents the IEnumTfLatticeElements COM interface.
type IEnumTfLatticeElements struct {
	IUnknown
}

// IEnumTfLatticeElementsVtbl COM interface vtable.
type IEnumTfLatticeElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfLatticeElements returns the IEnumTfLatticeElements vtable.
func (obj *IEnumTfLatticeElements) IEnumTfLatticeElements() *IEnumTfLatticeElementsVtbl {
	return (*IEnumTfLatticeElementsVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfLatticeElements) Clone(enum **IEnumTfLatticeElements) error {
	return result(call1(
		uintptr(obj.IEnumTfLatticeElements().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfLatticeElements) Next(count uint32, items *TFLattElement, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfLatticeElements().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfLatticeElements) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfLatticeElements().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfLatticeElements) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfLatticeElements().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
