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
