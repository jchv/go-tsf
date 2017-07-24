package tsf

// IEnumTfCandidates represents the IEnumTfCandidates COM interface.
type IEnumTfCandidates struct {
	IUnknown
}

// IEnumTfCandidatesVtbl COM interface vtable.
type IEnumTfCandidatesVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
