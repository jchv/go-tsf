package tsf

// IEnumTfDisplayAttributeInfo represents the IEnumTfDisplayAttributeInfo COM interface.
type IEnumTfDisplayAttributeInfo struct {
	IUnknown
}

// IEnumTfDisplayAttributeInfoVtbl COM interface vtable.
type IEnumTfDisplayAttributeInfoVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
