package tsf

// IEnumTfProperties represents the IEnumTfProperties COM interface.
type IEnumTfProperties struct {
	IUnknown
}

// IEnumTfPropertiesVtbl COM interface vtable.
type IEnumTfPropertiesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
