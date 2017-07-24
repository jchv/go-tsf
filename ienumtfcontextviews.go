package tsf

// IEnumTfContextViews represents the IEnumTfContextViews COM interface.
type IEnumTfContextViews struct {
	IUnknown
}

// IEnumTfContextViewsVtbl COM interface vtable.
type IEnumTfContextViewsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
