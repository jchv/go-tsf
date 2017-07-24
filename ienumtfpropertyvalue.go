package tsf

// IEnumTfPropertyValue represents the IEnumTfPropertyValue COM interface.
type IEnumTfPropertyValue struct {
	IUnknown
}

// IEnumTfPropertyValueVtbl COM interface vtable.
type IEnumTfPropertyValueVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
