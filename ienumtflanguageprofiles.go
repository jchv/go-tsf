package tsf

import "unsafe"

// IEnumTfLanguageProfiles represents the IEnumTfLanguageProfiles COM interface.
type IEnumTfLanguageProfiles struct {
	IUnknown
}

// IEnumTfLanguageProfilesVtbl COM interface vtable.
type IEnumTfLanguageProfilesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}

// IEnumTfLanguageProfiles returns the IEnumTfLanguageProfiles vtable.
func (obj *IEnumTfLanguageProfiles) IEnumTfLanguageProfiles() *IEnumTfLanguageProfilesVtbl {
	return (*IEnumTfLanguageProfilesVtbl)(unsafe.Pointer(obj.Vtbl))
}

// Clone creates a copy of the enumerator object.
func (obj *IEnumTfLanguageProfiles) Clone(enum **IEnumTfLanguageProfiles) error {
	return result(call1(
		uintptr(obj.IEnumTfLanguageProfiles().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(enum)),
	))
}

// Next obtains, from the current position, the specified number of elements
// in the enumeration sequence.
func (obj *IEnumTfLanguageProfiles) Next(count uint32, items *TFLanguageProfile, fetched *uint32) error {
	return result(call3(
		uintptr(obj.IEnumTfLanguageProfiles().Clone),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
		uintptr(unsafe.Pointer(items)),
		uintptr(unsafe.Pointer(fetched)),
	))
}

// Reset resets the enumerator object by moving the current position to the
// beginning of the enumeration sequence.
func (obj *IEnumTfLanguageProfiles) Reset() error {
	return result(call0(
		uintptr(obj.IEnumTfLanguageProfiles().Reset),
		uintptr(unsafe.Pointer(obj)),
	))
}

// Skip moves the current position forward in the enumeration sequence by the
// specified number of elements.
func (obj *IEnumTfLanguageProfiles) Skip(count uint32) error {
	return result(call1(
		uintptr(obj.IEnumTfLanguageProfiles().Skip),
		uintptr(unsafe.Pointer(obj)),
		uintptr(count),
	))
}
