package tsf

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
