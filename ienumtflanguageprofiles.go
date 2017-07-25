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
