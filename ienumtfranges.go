package tsf

import "unsafe"

// IEnumTfRanges represents the IEnumTfRanges COM interface.
type IEnumTfRanges struct {
	IUnknown
}

// IEnumTfRangesVtbl COM interface vtable.
type IEnumTfRangesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfRanges returns the IEnumTfRanges vtable.
func (obj *IEnumTfRanges) IEnumTfRanges() *IEnumTfRangesVtbl {
	return (*IEnumTfRangesVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfRanges) Clone(enum **IEnumTfRanges) error {
	return result(call1(
		uintptr(obj.IEnumTfRanges().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfRanges) Next(count uint32, items **ITfRange, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfRanges().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfRanges) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfRanges().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfRanges) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfRanges().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
