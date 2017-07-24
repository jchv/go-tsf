package tsf

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
