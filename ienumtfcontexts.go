package tsf

// IEnumTfContexts represents the IEnumTfContexts COM interface.
type IEnumTfContexts struct {
	IUnknown
}

// IEnumTfContextsVtbl COM interface vtable.
type IEnumTfContextsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
