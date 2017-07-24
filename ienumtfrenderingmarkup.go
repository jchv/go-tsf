package tsf

// IEnumTfRenderingMarkup represents the IEnumTfRenderingMarkup COM interface.
type IEnumTfRenderingMarkup struct {
	IUnknown
}

// IEnumTfRenderingMarkupVtbl COM interface vtable.
type IEnumTfRenderingMarkupVtbl struct {
	IUnknownVtbl
	Clone uintptr
	Next  uintptr
	Reset uintptr
	Skip  uintptr
}
