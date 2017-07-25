package tsf

import "unsafe"

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

// IEnumTfDocumentMgrs returns the IEnumTfDocumentMgrs vtable.
func (obj *IEnumTfDocumentMgrs) IEnumTfDocumentMgrs() *IEnumTfDocumentMgrsVtbl {
	return (*IEnumTfDocumentMgrsVtbl)(unsafe.Pointer(obj.Vtbl))
}
