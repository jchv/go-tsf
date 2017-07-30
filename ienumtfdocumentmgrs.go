package tsf

import "unsafe"

// IEnumTfDocumentMgrs represents the IEnumTfDocumentMgrs COM interface.
type IEnumTfDocumentMgrs struct {
	IUnknown
}

// IEnumTfDocumentMgrsVtbl COM interface vtable.
type IEnumTfDocumentMgrsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfDocumentMgrs returns the IEnumTfDocumentMgrs vtable.
func (obj *IEnumTfDocumentMgrs) IEnumTfDocumentMgrs() *IEnumTfDocumentMgrsVtbl {
	return (*IEnumTfDocumentMgrsVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfDocumentMgrs) Clone(enum **IEnumTfDocumentMgrs) error {
	return result(call1(
		uintptr(obj.IEnumTfDocumentMgrs().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfDocumentMgrs) Next(count uint32, items **ITfDocumentMgr, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfDocumentMgrs().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfDocumentMgrs) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfDocumentMgrs().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfDocumentMgrs) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfDocumentMgrs().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
