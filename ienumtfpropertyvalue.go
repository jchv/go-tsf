package tsf

import "unsafe"

// IEnumTfPropertyValue represents the IEnumTfPropertyValue COM interface.
type IEnumTfPropertyValue struct {
	IUnknown
}

// IEnumTfPropertyValueVtbl COM interface vtable.
type IEnumTfPropertyValueVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfPropertyValue returns the IEnumTfPropertyValue vtable.
func (obj *IEnumTfPropertyValue) IEnumTfPropertyValue() *IEnumTfPropertyValueVtbl {
	return (*IEnumTfPropertyValueVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfPropertyValue) Clone(enum **IEnumTfPropertyValue) error {
	return result(call1(
		uintptr(obj.IEnumTfPropertyValue().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfPropertyValue) Next(count uint32, items *TFPropertyVal, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfPropertyValue().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfPropertyValue) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfPropertyValue().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfPropertyValue) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfPropertyValue().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
