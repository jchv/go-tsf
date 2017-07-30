package tsf

import "unsafe"

// IEnumTfUIElements represents the IEnumTfUIElements COM interface.
type IEnumTfUIElements struct {
	IUnknown
}

// IEnumTfUIElementsVtbl COM interface vtable.
type IEnumTfUIElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfUIElements returns the IEnumTfUIElements vtable.
func (obj *IEnumTfUIElements) IEnumTfUIElements() *IEnumTfUIElementsVtbl {
	return (*IEnumTfUIElementsVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfUIElements) Clone(enum **IEnumTfUIElements) error {
	return result(call1(
		uintptr(obj.IEnumTfUIElements().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfUIElements) Next(count uint32, items **ITfUIElement, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfUIElements().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfUIElements) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfUIElements().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfUIElements) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfUIElements().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
