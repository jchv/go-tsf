package tsf

// IEnumTfLatticeElements represents the IEnumTfLatticeElements COM interface.
type IEnumTfLatticeElements struct {
	IUnknown
}

// IEnumTfLatticeElementsVtbl COM interface vtable.
type IEnumTfLatticeElementsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
