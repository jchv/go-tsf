package tsf

// IEnumTfDocumentMgrs represents the IEnumTfDocumentMgrs COM interface.
type IEnumTfDocumentMgrs struct {
	IUnknown
}

// IEnumTfDocumentMgrsVtbl COM interface vtable.
type IEnumTfDocumentMgrsVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
