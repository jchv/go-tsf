package tsf

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
