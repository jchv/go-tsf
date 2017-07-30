package tsf

import "unsafe"

// IEnumTfInputProcessorProfiles represents the IEnumTfInputProcessorProfiles COM interface.
type IEnumTfInputProcessorProfiles struct {
	IUnknown
}

// IEnumTfInputProcessorProfilesVtbl COM interface vtable.
type IEnumTfInputProcessorProfilesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfInputProcessorProfiles returns the IEnumTfInputProcessorProfiles vtable.
func (obj *IEnumTfInputProcessorProfiles) IEnumTfInputProcessorProfiles() *IEnumTfInputProcessorProfilesVtbl {
	return (*IEnumTfInputProcessorProfilesVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfInputProcessorProfiles) Clone(enum **IEnumTfInputProcessorProfiles) error {
	return result(call1(
		uintptr(obj.IEnumTfInputProcessorProfiles().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfInputProcessorProfiles) Next(count uint32, items *TFInputProcessorProfile, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfInputProcessorProfiles().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfInputProcessorProfiles) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfInputProcessorProfiles().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfInputProcessorProfiles) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfInputProcessorProfiles().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
