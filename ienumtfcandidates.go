package tsf

import "unsafe"

// IEnumTfCandidates represents the IEnumTfCandidates COM interface.
type IEnumTfCandidates struct {
	IUnknown
}

// IEnumTfCandidatesVtbl COM interface vtable.
type IEnumTfCandidatesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfCandidates returns the IEnumTfCandidates vtable.
func (obj *IEnumTfCandidates) IEnumTfCandidates() *IEnumTfCandidatesVtbl {
	return (*IEnumTfCandidatesVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfCandidates) Clone(enum **IEnumTfCandidates) error {
	return result(call1(
		uintptr(obj.IEnumTfCandidates().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfCandidates) Next(count uint32, items **ITfCandidateString, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfCandidates().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfCandidates) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfCandidates().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfCandidates) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfCandidates().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
