package tsf

import "unsafe"

// IEnumTfRenderingMarkup represents the IEnumTfRenderingMarkup COM interface.
type IEnumTfRenderingMarkup struct {
	IUnknown
}

// IEnumTfRenderingMarkupVtbl COM interface vtable.
type IEnumTfRenderingMarkupVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfRenderingMarkup returns the IEnumTfRenderingMarkup vtable.
func (obj *IEnumTfRenderingMarkup) IEnumTfRenderingMarkup() *IEnumTfRenderingMarkupVtbl {
	return (*IEnumTfRenderingMarkupVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfRenderingMarkup) Clone(enum **IEnumTfRenderingMarkup) error {
	return result(call1(
		uintptr(obj.IEnumTfRenderingMarkup().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfRenderingMarkup) Next(count uint32, items *TFRenderingMarkup, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfRenderingMarkup().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfRenderingMarkup) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfRenderingMarkup().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfRenderingMarkup) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfRenderingMarkup().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
