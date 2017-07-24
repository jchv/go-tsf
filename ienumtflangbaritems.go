package tsf

// IEnumTfLangBarItems represents the IEnumTfLangBarItems COM interface.
type IEnumTfLangBarItems struct {
	IUnknown
}

// IEnumTfLangBarItemsVtbl COM interface vtable.
type IEnumTfLangBarItemsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
