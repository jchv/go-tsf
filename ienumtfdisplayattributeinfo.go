package tsf

import "unsafe"

// IEnumTfDisplayAttributeInfo represents the IEnumTfDisplayAttributeInfo COM interface.
type IEnumTfDisplayAttributeInfo struct {
	IUnknown
}

// IEnumTfDisplayAttributeInfoVtbl COM interface vtable.
type IEnumTfDisplayAttributeInfoVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfDisplayAttributeInfo returns the IEnumTfDisplayAttributeInfo vtable.
func (obj *IEnumTfDisplayAttributeInfo) IEnumTfDisplayAttributeInfo() *IEnumTfDisplayAttributeInfoVtbl {
	return (*IEnumTfDisplayAttributeInfoVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfDisplayAttributeInfo) Clone(enum **IEnumTfDisplayAttributeInfo) error {
	return result(call1(
		uintptr(obj.IEnumTfDisplayAttributeInfo().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfDisplayAttributeInfo) Next(count uint32, items **ITfDisplayAttributeInfo, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfDisplayAttributeInfo().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfDisplayAttributeInfo) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfDisplayAttributeInfo().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfDisplayAttributeInfo) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfDisplayAttributeInfo().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
