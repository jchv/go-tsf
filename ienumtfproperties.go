package tsf

import "unsafe"

// IEnumTfProperties represents the IEnumTfProperties COM interface.
type IEnumTfProperties struct {
	IUnknown
}

// IEnumTfPropertiesVtbl COM interface vtable.
type IEnumTfPropertiesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfProperties returns the IEnumTfProperties vtable.
func (obj *IEnumTfProperties) IEnumTfProperties() *IEnumTfPropertiesVtbl {
	return (*IEnumTfPropertiesVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfProperties) Clone(enum **IEnumTfProperties) error {
	return result(call1(
		uintptr(obj.IEnumTfProperties().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfProperties) Next(count uint32, items **ITfProperty, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfProperties().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfProperties) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfProperties().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfProperties) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfProperties().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
