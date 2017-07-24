package tsf

// IEnumITfCompositionView represents the IEnumITfCompositionView COM interface.
type IEnumITfCompositionView struct {
	IUnknown
}

// IEnumITfCompositionViewVtbl COM interface vtable.
type IEnumITfCompositionViewVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
