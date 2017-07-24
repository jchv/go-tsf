package tsf

// EnumTFUIElements represents the EnumTFUIElements COM interface.
type EnumTFUIElements struct {
	IUnknown
}

// EnumTFUIElementsVtbl COM interface vtable.
type EnumTFUIElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
