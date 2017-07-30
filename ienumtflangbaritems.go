package tsf

import "unsafe"

// IEnumTfLangBarItems represents the IEnumTfLangBarItems COM interface.
type IEnumTfLangBarItems struct {
	IUnknown
}

// IEnumTfLangBarItemsVtbl COM interface vtable.
type IEnumTfLangBarItemsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfLangBarItems returns the IEnumTfLangBarItems vtable.
func (obj *IEnumTfLangBarItems) IEnumTfLangBarItems() *IEnumTfLangBarItemsVtbl {
	return (*IEnumTfLangBarItemsVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfLangBarItems) Clone(enum **IEnumTfLangBarItems) error {
	return result(call1(
		uintptr(obj.IEnumTfLangBarItems().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfLangBarItems) Next(count uint32, items **ITfLangBarItem, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfLangBarItems().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfLangBarItems) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfLangBarItems().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfLangBarItems) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfLangBarItems().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
